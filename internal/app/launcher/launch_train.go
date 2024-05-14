package launcher

import (
	bytes2 "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/hse-experiments-platform/launcher/internal/pkg/domain"
	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *launcherService) LaunchTrain(ctx context.Context, req *pb.LaunchTrainRequest) (*pb.LaunchTrainResponse, error) {
	if req.GetDatasetSettings() == nil || req.GetDatasetSettings().GetDatasetID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "datasetID must be not 0")
	}

	userID, err := s.checkDatasetAccess(ctx, int64(req.GetDatasetSettings().GetDatasetID()))
	if err != nil {
		return nil, err
	}

	name, err := s.commonDB.CheckModelExists(ctx, int64(req.GetModelSettings().GetModelID()))
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "model not found")
	} else if err != nil {
		return nil, fmt.Errorf("s.commonDB.CheckModelExists: %w", err)
	}

	metrics, err := s.commonDB.GetModelMetrics(ctx, int64(req.GetModelSettings().GetModelID()))
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.GetModelMetrics: %w", err)
	}

	hyperparamers, err := s.commonDB.GetModelHyperparameters(ctx, int64(req.GetModelSettings().GetModelID()))
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.GetModelHyperparameters: %w", err)
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("input json.Marshal: %w", err)
	}

	launchID, err := s.commonDB.CreateLaunch(ctx, db.CreateLaunchParams{
		LaunchType:   pb.LaunchType_LaunchTypeTrain.String(),
		UserID:       userID,
		Name:         req.GetLaunchInfo().GetName(),
		Description:  req.GetLaunchInfo().GetDescription(),
		LaunchStatus: pb.LaunchStatus_LaunchStatusNotStarted.String(),
		Input:        bytes,
	})
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.CreateLaunch: %w", err)
	}

	modelID, err := s.commonDB.CreateTrainedModel(ctx, db.CreateTrainedModelParams{
		Name:           pgtype.Text{String: req.GetLaunchInfo().GetName(), Valid: true},
		Description:    pgtype.Text{String: req.GetLaunchInfo().GetDescription(), Valid: true},
		BaseModelID:    pgtype.Int8{Int64: int64(req.GetModelSettings().GetModelID()), Valid: true},
		TrainDatasetID: pgtype.Int8{Int64: int64(req.GetDatasetSettings().GetDatasetID()), Valid: true},
		LaunchID:       pgtype.Int8{Int64: launchID, Valid: true},
		TargetColumn:   pgtype.Text{String: req.GetDatasetSettings().GetTargetColumn(), Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.CreateTrainedModel: %w", err)
	}
	log.Debug().Int64("model_id", modelID).Msg("trained model created")

	if err = s.launcher.Launch(launchID, s.trainLauncher(userID, req, name.String, metrics, hyperparamers)); err != nil {
		return nil, fmt.Errorf("s.launcher.Launch: %w", err)
	}

	return &pb.LaunchTrainResponse{
		LaunchID: launchID,
	}, nil
}

func (s *launcherService) trainLauncher(userID int64, req *pb.LaunchTrainRequest, name string, metrics []string, hyperparamers []db.GetModelHyperparametersRow) func(context.Context, int64) ([]byte, error) {
	return func(ctx context.Context, launchID int64) ([]byte, error) {
		body, err := makeTrainRequest(launchID, userID, req, name, metrics, hyperparamers)
		if err != nil {
			return nil, fmt.Errorf("makeUploadRequest: %w", err)
		}

		var bytes []byte
		var errM error

		resp, err := s.client.SendTrainTask(ctx, body)
		if resp != "" {
			if bytes, errM = json.Marshal(resp); errM != nil {
				log.Error().Err(errM).Msg("json.Marshal")
			}
		}
		if err != nil {
			return bytes, fmt.Errorf("s.client.SendDatasetSetTypesTask: %w", err)
		}

		log.Debug().Str("resp", fmt.Sprint(resp)).Msg("response")

		return bytes, nil
	}
}

func makeTrainRequest(launchID, userID int64, req *pb.LaunchTrainRequest, name string, metrics []string, hyperparamers []db.GetModelHyperparametersRow) (io.Reader, error) {
	type trainParams struct {
		UseCV               bool    `json:"use_cv"`
		CVChunks            int     `json:"cv_chunks"`
		TrainTestSplitRatio float64 `json:"train_test_split_ratio"`
		Seed                int     `json:"seed"`
	}

	var body struct {
		UserID          string                 `json:"user_id"`
		DatasetID       string                 `json:"dataset_id"`
		LaunchID        string                 `json:"launch_id"`
		TargetCol       string                 `json:"target_col"`
		TrainModelName  string                 `json:"train_model_name"`
		Hyperparameters map[string]interface{} `json:"hyperparameters"`
		Metrics         []string               `json:"metrics"`
		TrainParams     trainParams            `json:"train_params"`
	}

	body.LaunchID = fmt.Sprintf("%d", launchID)
	body.UserID = domain.GetBucketName(userID)
	body.DatasetID = domain.GetObjectName(int64(req.GetDatasetSettings().GetDatasetID()))
	body.TargetCol = req.GetDatasetSettings().GetTargetColumn()
	body.TrainModelName = name
	body.Metrics = metrics
	body.TrainParams = trainParams{
		UseCV:               req.GetUseCv(),
		CVChunks:            int(req.GetCvChunks()),
		TrainTestSplitRatio: req.GetDatasetSettings().GetTrainTestSplit(),
		Seed:                int(req.GetRandomSeed()),
	}
	body.Hyperparameters = make(map[string]interface{}, len(hyperparamers))
	for _, hp := range hyperparamers {
		if req.GetModelSettings().GetHyperparameterValues() != nil {
			if v, ok := req.ModelSettings.HyperparameterValues[uint64(hp.ID)]; ok {
				body.Hyperparameters[hp.Name] = v
			}
			//else {
			//    body.Hyperparameters[hp.Name] = hp.DefaultValue
			//}
		}
		//else {
		//    body.Hyperparameters[hp.Name] = hp.DefaultValue
		//}
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	log.Debug().Str("body", string(bytes)).Msg("request")

	return bytes2.NewReader(bytes), nil
}

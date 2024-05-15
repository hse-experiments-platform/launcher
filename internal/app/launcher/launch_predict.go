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
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *launcherService) LaunchPredict(ctx context.Context, req *pb.LaunchPredictRequest) (*pb.LaunchPredictResponse, error) {
	if req.GetDatasetID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "datasetID must be not 0")
	}

	userID, err := s.checkDatasetAccess(ctx, req.GetDatasetID())
	if err != nil {
		return nil, err
	}

	runIDJson, err := s.commonDB.GetTrainedModelRunID(ctx, req.GetTrainedModelID())
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "trained model not found")
	} else if err != nil {
		return nil, fmt.Errorf("s.commonDB.GetTrainedModelRunID: %w", err)
	}

	var runID string
	if err = json.Unmarshal(runIDJson, &runID); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("input json.Marshal: %w", err)
	}
	launchID, err := s.commonDB.CreateLaunch(ctx, db.CreateLaunchParams{
		LaunchType:   pb.LaunchType_LaunchTypePredict.String(),
		UserID:       userID,
		Name:         req.GetLaunchInfo().GetName(),
		Description:  req.GetLaunchInfo().GetDescription(),
		LaunchStatus: pb.LaunchStatus_LaunchStatusNotStarted.String(),
		Input:        bytes,
	})
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.CreateLaunch: %w", err)
	}

	if err = s.launcher.Launch(launchID, s.predictLauncher(userID, req, runID)); err != nil {
		return nil, fmt.Errorf("s.launcher.Launch: %w", err)
	}

	return &pb.LaunchPredictResponse{
		LaunchID: launchID,
	}, nil
}

func (s *launcherService) predictLauncher(userID int64, req *pb.LaunchPredictRequest, runID string) func(context.Context, int64) ([]byte, error) {
	return func(ctx context.Context, launchID int64) ([]byte, error) {
		body, err := makePredictRequest(launchID, userID, req, runID)
		if err != nil {
			return nil, fmt.Errorf("makeUploadRequest: %w", err)
		}

		var bytes []byte
		var errM error

		resp, err := s.client.SendPredictTask(ctx, body)
		if resp != nil {
			if bytes, errM = json.Marshal(resp); errM != nil {
				log.Error().Err(errM).Msg("json.Marshal")
			}
		}
		if err != nil {
			return bytes, fmt.Errorf("s.client.SendPredictTask: %w", err)
		}

		if resp["object_name"] == nil {
			return bytes, fmt.Errorf("object_name is nil")
		}

		log.Debug().Str("resp", fmt.Sprint(resp)).Msg("response")

		return bytes, nil
	}
}

func makePredictRequest(launchID, userID int64, req *pb.LaunchPredictRequest, runID string) (io.Reader, error) {
	var body struct {
		UserID        string `json:"user_id"`
		DatasetID     string `json:"dataset_id"`
		LaunchID      string `json:"launch_id"`
		TrainingRunID string `json:"training_run_id"`
	}

	body.UserID = domain.GetBucketName(userID)
	body.DatasetID = fmt.Sprint(req.GetDatasetID())
	body.LaunchID = fmt.Sprint(launchID)
	body.TrainingRunID = runID

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	log.Debug().Str("body", string(bytes)).Msg("request")

	return bytes2.NewReader(bytes), nil
}

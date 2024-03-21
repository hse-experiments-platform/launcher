package launcher

import (
	"context"
	"fmt"

	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *launcherService) StartTraining(ctx context.Context, req *pb.StartTrainingRequest) (*pb.StartTrainingResponse, error) {
	resp := &pb.StartTrainingResponse{}
	launchID := int64(0)
	modelID := int64(0)

	err := pgx.BeginTxFunc(ctx, s.commonDBConn, pgx.TxOptions{
		AccessMode: pgx.ReadWrite,
	}, func(tx pgx.Tx) (err error) {
		defer func() {
			if err != nil {
				_ = tx.Rollback(ctx)
			}
		}()

		dbtx := s.commonDB.WithTx(tx)
		launchID, err = dbtx.CreateLaunch(ctx, db.CreateLaunchParams{
			LaunchType:  db.LaunchTypeTrain,
			Name:        req.GetName(),
			Description: req.GetDescription(),
		})
		if err != nil {
			return fmt.Errorf("dbtx.CreateLaunch: %w", err)
		}

		modelID, err = dbtx.CreateTrainedModel(ctx, db.CreateTrainedModelParams{
			Name:                req.GetName(),
			Description:         req.GetDescription(),
			ModelID:             int64(req.GetModelSettings().GetModelID()),
			ModelTrainingStatus: db.ModelTrainingStatusInProgress,
			TrainingDatasetID:   int64(req.GetDatasetSettings().GetDatasetID()),
			TargetColumn:        req.GetDatasetSettings().GetTargetColumn(),
			LaunchID:            launchID,
		})
		if err != nil {
			return fmt.Errorf("dbtx.CreateTrainedModel: %w", err)
		}

		return nil
	})
	if err != nil {
		if err := s.failStartTraining(ctx, modelID, launchID, err); err != nil {
			return nil, fmt.Errorf("s.failStartTraining: %w", err)
		}

		return nil, fmt.Errorf("pgx.BeginTxFunc: %w", err)
	}

	params := db.CreateTrainingHyperparametersParams{TrainModelID: modelID}
	for id, v := range req.GetModelSettings().GetHyperparameterValues() {
		params.HyperparameterIds = append(params.HyperparameterIds, int64(id))
		params.Values = append(params.Values, []byte(v))
	}
	if err := s.commonDB.CreateTrainingHyperparameters(ctx, params); err != nil {
		if err := s.failStartTraining(ctx, modelID, launchID, err); err != nil {
			return nil, fmt.Errorf("s.failStartTraining: %w", err)
		}

		return nil, fmt.Errorf("dbtx.CreateTrainingHyperparameters: %w", err)
	}

	resp = &pb.StartTrainingResponse{
		TrainedModelID: uint64(modelID),
		LaunchID:       uint64(launchID),
	}
	return resp, nil
}

func (s *launcherService) failStartTraining(ctx context.Context, modelID int64, launchID int64, err error) error {
	if modelID != 0 {
		if errS := s.commonDB.UpdateTrainedModelStatus(ctx, db.UpdateTrainedModelStatusParams{
			ID:                  modelID,
			TrainError:          pgtype.Text{String: err.Error()},
			ModelTrainingStatus: db.ModelTrainingStatusError,
		}); errS != nil {
			return fmt.Errorf("cannot set erro: %s, error: %w", errS.Error(), err)
		}
	}
	if launchID != 0 {
		if errS := s.commonDB.UpdateLaunchStatus(ctx, db.UpdateLaunchStatusParams{
			ID:          modelID,
			LaunchError: pgtype.Text{String: err.Error()},
		}); errS != nil {
			return fmt.Errorf("cannot set erro: %s, error: %w", errS.Error(), err)
		}
	}
	return nil
}

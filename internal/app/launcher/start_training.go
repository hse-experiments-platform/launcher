package launcher

import (
	"context"
	"fmt"

	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
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
		if modelID != 0 {
			if err := s.commonDB.UpdateTrainedModelStatus(ctx, db.UpdateTrainedModelStatusParams{
				ID:                  modelID,
				TrainError:          pgtype.Text{String: err.Error()},
				ModelTrainingStatus: db.ModelTrainingStatusError,
			}); err != nil {
				log.Error().Err(err).Msg("cannot set trained model error")
			}
		}
		if launchID != 0 {
			if err := s.commonDB.UpdateLaunchStatus(ctx, db.UpdateLaunchStatusParams{
				ID:          modelID,
				LaunchError: pgtype.Text{String: err.Error()},
			}); err != nil {
				log.Error().Err(err).Msg("cannot set launch error")
			}
		}
		return nil, fmt.Errorf("pgx.BeginTxFunc: %w", err)
	}

	params := db.CreateTrainingHyperparametersParams{TrainModelID: modelID}
	for id, v := range req.GetModelSettings().GetHyperparameterValues() {
		params.HyperparameterIds = append(params.HyperparameterIds, int64(id))
		params.Values = append(params.Values, []byte(v))
	}
	if err1 := s.commonDB.CreateTrainingHyperparameters(ctx, params); err != nil {
		return nil, fmt.Errorf("dbtx.CreateTrainingHyperparameters: %w", err1)
	}

	resp = &pb.StartTrainingResponse{
		TrainedModelID: uint64(modelID),
		LaunchID:       uint64(launchID),
	}
	return resp, nil
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CheckDatasetStatus(ctx context.Context, arg CheckDatasetStatusParams) (bool, error)
	CheckLaunchStatus(ctx context.Context, arg CheckLaunchStatusParams) (bool, error)
	CheckModelExists(ctx context.Context, id int64) (pgtype.Text, error)
	CreateLaunch(ctx context.Context, arg CreateLaunchParams) (int64, error)
	CreateTrainedModel(ctx context.Context, arg CreateTrainedModelParams) (int64, error)
	FinishLaunch(ctx context.Context, arg FinishLaunchParams) error
	GetDatasetCreator(ctx context.Context, id int64) (int64, error)
	GetLaunch(ctx context.Context, id int64) (GetLaunchRow, error)
	GetLaunchType(ctx context.Context, id int64) (string, error)
	GetLaunches(ctx context.Context, arg GetLaunchesParams) ([]GetLaunchesRow, error)
	GetModelHyperparameters(ctx context.Context, id int64) ([]GetModelHyperparametersRow, error)
	GetModelMetrics(ctx context.Context, id int64) ([]string, error)
	GetTrainedModelIDByLaunchID(ctx context.Context, launchID pgtype.Int8) (int64, error)
	UpdateLaunchStatus(ctx context.Context, arg UpdateLaunchStatusParams) error
}

var _ Querier = (*Queries)(nil)

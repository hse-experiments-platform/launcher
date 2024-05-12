package launcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hse-experiments-platform/launcher/internal/pkg/core/launch"
	"github.com/hse-experiments-platform/launcher/internal/pkg/interactions/workers"
	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/hse-experiments-platform/library/pkg/utils/token"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.LauncherServiceServer = (*launcherService)(nil)

type launcherService struct {
	pb.UnimplementedLauncherServiceServer
	commonDBConn *pgxpool.Pool
	maker        token.Maker
	client       workers.WorkersClient

	commonDB *db.Queries
	launcher launch.Launcher
}

func NewService(commonDBConn *pgxpool.Pool, maker token.Maker, client workers.WorkersClient) *launcherService {
	return &launcherService{
		commonDBConn: commonDBConn,
		maker:        maker,
		client:       client,

		commonDB: db.New(commonDBConn),
		launcher: launch.NewLauncher(commonDBConn),
	}
}

func getUserID(ctx context.Context) (int64, error) {
	var userID int64
	userID, ok := ctx.Value(token.UserIDContextKey).(int64)
	if !ok {
		log.Error().Msgf("invalid userID context key type: %T", ctx.Value(token.UserIDContextKey))
		return 0, status.New(codes.Internal, "internal error").Err()
	}

	return userID, nil
}

func (s *launcherService) checkDatasetAccess(ctx context.Context, datasetID int64) (int64, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return 0, err
	}

	creatorID, err := s.commonDB.GetDatasetCreator(ctx, datasetID)
	if errors.Is(err, pgx.ErrNoRows) {
		return userID, status.Error(codes.NotFound, "dataset not found")
	} else if err != nil {
		return userID, fmt.Errorf("d.commonDB.GetDatasetCreator: %w", err)
	}

	if creatorID != userID {
		return userID, status.Error(codes.PermissionDenied, "cannot get other user's dataset rows")
	}

	return userID, nil
}

func parseLaunchInputOutput[InputT any, OutputT any](ctx context.Context, s *launcherService, launchID int64) (db.GetLaunchRow, InputT, OutputT, error) {
	var input InputT
	var output OutputT

	launch, err := s.commonDB.GetLaunch(ctx, launchID)
	if err != nil {
		return launch, input, output, fmt.Errorf("s.commonDB.GetLaunch: %w", err)
	}

	err = json.Unmarshal(launch.Input, &input)
	if err != nil {
		return launch, input, output, fmt.Errorf("json.Unmarshal input: %w", err)
	}

	if launch.LaunchStatus == pb.LaunchStatus_LaunchStatusSuccess.String() {
		err = json.Unmarshal(launch.Output, &output)
		if err != nil {
			return launch, input, output, fmt.Errorf("json.Unmarshal output: %w", err)
		}
	}

	return launch, input, output, nil
}

package launcher

import (
	"context"
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

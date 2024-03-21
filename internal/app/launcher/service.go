package launcher

import (
	"context"

	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/hse-experiments-platform/library/pkg/utils/token"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.LauncherServiceServer = (*launcherService)(nil)

type launcherService struct {
	pb.UnimplementedLauncherServiceServer
	commonDBConn *pgx.Conn
	maker        token.Maker

	commonDB *db.Queries
}

func NewService(commonDBConn *pgx.Conn, maker token.Maker) *launcherService {
	return &launcherService{
		commonDBConn: commonDBConn,
		maker:        maker,

		commonDB: db.New(commonDBConn),
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

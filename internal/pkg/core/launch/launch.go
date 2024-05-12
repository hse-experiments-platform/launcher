package launch

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/hse-experiments-platform/library/pkg/utils/errors/revertable"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Launcher interface {
	Launch(launchID int64, launchFunc func(context.Context, int64) ([]byte, error)) error
}

var _ Launcher = (*launcher)(nil)

type launcher struct {
	commonDBConn *pgxpool.Pool
	commonDB     *db.Queries
}

func NewLauncher(commonDBConn *pgxpool.Pool) *launcher {
	return &launcher{
		commonDBConn: commonDBConn,
		commonDB:     db.New(commonDBConn),
	}
}

func (l *launcher) Launch(launchID int64, launchFunc func(context.Context, int64) ([]byte, error)) error {
	log.Debug().Int64("launch_id", launchID).Msgf("starting launch %v", launchID)
	ctx, c := context.WithTimeout(context.Background(), time.Minute*30)

	go func() {
		defer c()
		l.launch(ctx, launchID, launchFunc)
	}()

	return nil
}

func (l *launcher) launch(ctx context.Context, launchID int64, launchFunc func(context.Context, int64) ([]byte, error)) {
	var err error
	var bytes []byte

	// depending on the error, we either finish the launch with error or success
	defer func() {
		if err != nil {
			err = fmt.Errorf("launchExecute: %w", err)
			l.onLaunchFailed(ctx, launchID, err, bytes)
		} else {
			l.onLaunchSucceeded(ctx, launchID, bytes)
		}
	}()
	// convert panic to our err
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	bytes, err = launchFunc(ctx, launchID)
}

func (l *launcher) onLaunchFailed(ctx context.Context, launchID int64, err error, bytes []byte) {
	var revErr *revertable.RevertableError
	if !errors.As(err, &revErr) {
		revErr = revertable.NewRevertable(err, "internal error")
	}

	log.Error().Err(err).Msg("launch failed")
	if err := l.commonDB.FinishLaunch(ctx, db.FinishLaunchParams{
		ID:           launchID,
		LaunchStatus: pb.LaunchStatus_LaunchStatusError.String(),
		LaunchError: pgtype.Text{
			String: revErr.GetReason(),
			Valid:  true,
		},
		Output: bytes,
	}); err != nil {
		log.Error().Err(err).AnErr("fail_err", revErr).Msg("failed to finish launch")
	}
}

func (l *launcher) onLaunchSucceeded(ctx context.Context, launchID int64, bytes []byte) {
	log.Debug().Int64("launch_id", launchID).Msg("launch succeeded")
	if err := l.commonDB.FinishLaunch(ctx, db.FinishLaunchParams{
		ID:           launchID,
		LaunchStatus: pb.LaunchStatus_LaunchStatusSuccess.String(),
		Output:       bytes,
	}); err != nil {
		log.Error().Err(err).Msg("failed to finish launch")
	}
}

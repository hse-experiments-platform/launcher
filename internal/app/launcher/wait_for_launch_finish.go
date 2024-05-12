package launcher

import (
	"context"
	"fmt"
	"time"

	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
)

func (s *launcherService) WaitForLaunchFinish(ctx context.Context, req *pb.WaitForLaunchFinishRequest) (*pb.WaitForLaunchFinishResponse, error) {
	ctx, c := context.WithTimeout(ctx, time.Minute*30)
	defer c()

	t := time.NewTicker(time.Millisecond * 250)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.C:
			finished, err := s.commonDB.CheckLaunchStatus(ctx, db.CheckLaunchStatusParams{
				ID:       req.LaunchID,
				Statuses: []string{pb.LaunchStatus_LaunchStatusError.String(), pb.LaunchStatus_LaunchStatusSuccess.String()},
			})
			if err != nil {
				return nil, fmt.Errorf("l.commonDB.CheckLaunchStatus: %w", err)
			}

			if finished {
				return &pb.WaitForLaunchFinishResponse{}, nil
			}
		}
	}

}

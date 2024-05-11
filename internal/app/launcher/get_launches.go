package launcher

import (
	"context"
	"fmt"

	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *launcherService) GetLaunches(ctx context.Context, req *pb.GetLaunchesRequest) (*pb.GetLaunchesResponse, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	launchType := ""
	if req.GetLaunchType() != pb.LaunchType_LaunchTypeUnknown {
		launchType = req.GetLaunchType().String()
	}

	launches, err := s.commonDB.GetLaunches(ctx, db.GetLaunchesParams{
		UserID: userID,
		Name: pgtype.Text{
			String: req.GetNameQuery(),
			Valid:  true,
		},
		LaunchType: launchType,
		Limit:      int64(req.GetLimit()),
		Offset:     int64(req.GetOffset()),
	})
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.GetLaunches: %w", err)
	}
	count := 0
	if len(launches) > 0 {
		count = int(launches[0].Count)
	}

	return &pb.GetLaunchesResponse{Launches: convertLaunches(launches), PageInfo: &pb.PageInfo{
		Offset: req.GetOffset(),
		Limit:  req.GetLimit(),
		Total:  uint64(count),
	}}, nil
}

func convertLaunches(launches []db.GetLaunchesRow) []*pb.LaunchInfo {
	res := make([]*pb.LaunchInfo, 0, len(launches))

	for _, launch := range launches {
		res = append(res, &pb.LaunchInfo{
			LaunchID:    uint64(launch.ID),
			LaunchType:  pb.LaunchType(pb.LaunchType_value[launch.LaunchType]),
			Name:        launch.Name,
			Description: launch.Description,
			Status:      pb.LaunchStatus(pb.LaunchStatus_value[launch.LaunchStatus]),
		})
	}

	return res
}

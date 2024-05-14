package launcher

import (
	"context"
	"fmt"

	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
)

func (s *launcherService) GetDatasetSetTypesLaunch(ctx context.Context, req *pb.GetDatasetSetTypesLaunchRequest) (*pb.GetDatasetSetTypesLaunchResponse, error) {
	if err := s.checkLaunch(ctx, req.GetLaunchID(), pb.LaunchType_LaunchTypeDatasetSetTypes); err != nil {
		return nil, err
	}

	launch, input, _, err := parseLaunchInputOutput[pb.LaunchDatasetSetTypesRequest, any](ctx, s, req.GetLaunchID())
	if err != nil {
		return nil, fmt.Errorf("parseLaunchInputOutput: %w", err)
	}

	return &pb.GetDatasetSetTypesLaunchResponse{
		LaunchInfo: &pb.LaunchInfo{
			LaunchID:    uint64(launch.ID),
			Name:        input.LaunchInfo.Name,
			Description: input.LaunchInfo.Description,
			LaunchType:  pb.LaunchType(pb.LaunchType_value[launch.LaunchType]),
			Status:      pb.LaunchStatus(pb.LaunchStatus_value[launch.LaunchStatus]),
		},
		Error:       launch.LaunchError.String,
		DatasetID:   input.DatasetID,
		ColumnTypes: input.ColumnTypes,
	}, nil

}

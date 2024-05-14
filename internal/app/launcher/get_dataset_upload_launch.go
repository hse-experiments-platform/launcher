package launcher

import (
	"context"
	"fmt"

	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
)

func (s *launcherService) GetDatasetUploadLaunch(ctx context.Context, req *pb.GetDatasetUploadLaunchRequest) (*pb.GetDatasetUploadLaunchResponse, error) {
	if err := s.checkLaunch(ctx, req.GetLaunchID(), pb.LaunchType_LaunchTypeDatasetUpload); err != nil {
		return nil, err
	}

	launch, input, output, err := parseLaunchInputOutput[pb.LaunchDatasetUploadRequest, []string](ctx, s, req.GetLaunchID())
	if err != nil {
		return nil, fmt.Errorf("parseLaunchInputOutput: %w", err)
	}

	cols := make(map[string]pb.ColumnType, len(output))
	for _, v := range output {
		cols[v] = pb.ColumnType_ColumnTypeUndefined
	}

	return &pb.GetDatasetUploadLaunchResponse{
		LaunchInfo: &pb.LaunchInfo{
			LaunchID:    uint64(launch.ID),
			Name:        input.LaunchInfo.Name,
			Description: input.LaunchInfo.Description,
			LaunchType:  pb.LaunchType(pb.LaunchType_value[launch.LaunchType]),
			Status:      pb.LaunchStatus(pb.LaunchStatus_value[launch.LaunchStatus]),
		},
		Error:       launch.LaunchError.String,
		DatasetID:   input.DatasetID,
		ColumnTypes: cols,
	}, nil
}

package launcher

import (
	"context"
	"fmt"

	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
)

func (s *launcherService) GetPredictLaunch(ctx context.Context, req *pb.GetPredictLaunchRequest) (*pb.GetPredictLaunchResponse, error) {
	if err := s.checkLaunch(ctx, req.GetLaunchID(), pb.LaunchType_LaunchTypePredict); err != nil {
		return nil, err
	}

	launch, input, _, err := parseLaunchInputOutput[pb.LaunchPredictRequest, map[string]any](ctx, s, req.GetLaunchID())
	if err != nil {
		return nil, fmt.Errorf("parseLaunchInputOutput: %w", err)
	}

	return &pb.GetPredictLaunchResponse{
		LaunchInfo: &pb.LaunchInfo{
			LaunchID:    uint64(launch.ID),
			Name:        input.LaunchInfo.Name,
			Description: input.LaunchInfo.Description,
			LaunchType:  pb.LaunchType(pb.LaunchType_value[launch.LaunchType]),
			Status:      pb.LaunchStatus(pb.LaunchStatus_value[launch.LaunchStatus]),
		},
		Error:          launch.LaunchError.String,
		DatasetID:      input.GetDatasetID(),
		TrainedModelID: input.GetTrainedModelID(),
	}, nil
}

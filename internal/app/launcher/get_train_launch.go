package launcher

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *launcherService) checkLaunch(ctx context.Context, launchID int64, launchType pb.LaunchType) error {
	if t, err := s.commonDB.GetLaunchType(ctx, launchID); err != nil && errors.Is(err, pgx.ErrNoRows) {
		return status.Error(codes.NotFound, "launch not found")
	} else if err != nil {
		return fmt.Errorf("s.commonDB.GetLaunchType: %w", err)
	} else if t != launchType.String() {
		return status.Errorf(codes.InvalidArgument, "launchID is not a %s launch", launchType.String())
	}

	return nil
}

func (s *launcherService) GetTrainLaunch(ctx context.Context, req *pb.GetTrainLaunchRequest) (*pb.GetTrainLaunchResponse, error) {
	if err := s.checkLaunch(ctx, req.GetLaunchID(), pb.LaunchType_LaunchTypeTrain); err != nil {
		return nil, err
	}

	launch, input, output, err := parseLaunchInputOutput[pb.LaunchTrainRequest, string](ctx, s, req.GetLaunchID())
	if err != nil {
		return nil, fmt.Errorf("parseLaunchInputOutput: %w", err)
	}

	modelID, err := s.commonDB.GetTrainedModelIDByLaunchID(ctx, pgtype.Int8{Int64: req.GetLaunchID(), Valid: true})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("s.commonDB.GetTrainedModelIDByLaunchID: %w", err)
	}

	return &pb.GetTrainLaunchResponse{
		LaunchInfo: &pb.LaunchInfo{
			LaunchID:    uint64(launch.ID),
			Name:        input.LaunchInfo.Name,
			Description: input.LaunchInfo.Description,
			LaunchType:  pb.LaunchType(pb.LaunchType_value[launch.LaunchType]),
			Status:      pb.LaunchStatus(pb.LaunchStatus_value[launch.LaunchStatus]),
		},
		Error:          launch.LaunchError.String,
		DatasetID:      int64(input.GetDatasetSettings().GetDatasetID()),
		MlflowRunID:    output,
		TrainedModelID: modelID,
	}, nil
}

package launcher

import (
	"context"
	"fmt"

	"github.com/hse-experiments-platform/launcher/internal/pkg/domain"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *launcherService) GetPredictionResultDownloadLink(ctx context.Context, req *pb.GetPredictionResultDownloadLinkRequest) (*pb.GetPredictionResultDownloadLinkResponse, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.checkLaunch(ctx, req.GetLaunchID(), pb.LaunchType_LaunchTypePredict); err != nil {
		return nil, err
	}

	_, _, output, err := parseLaunchInputOutput[pb.LaunchPredictRequest, map[string]any](ctx, s, req.GetLaunchID())
	if err != nil {
		return nil, fmt.Errorf("parseLaunchInputOutput: %w", err)
	}

	var objectName string
	var ok bool
	if output["object_name"] == nil {
		return nil, status.Error(codes.NotFound, "prediction result not found")
	} else if objectName, ok = output["object_name"].(string); !ok {
		return nil, status.Error(codes.NotFound, "prediction result not found")
	}

	url, err := s.s3Storage.GetObjectDownloadLink(ctx, domain.GetBucketName(userID), objectName)
	if err != nil {
		return nil, fmt.Errorf("s.s3Storage.GetObjectDownloadLink: %w", err)
	}

	return &pb.GetPredictionResultDownloadLinkResponse{
		DownloadLink: url.String(),
	}, nil
}

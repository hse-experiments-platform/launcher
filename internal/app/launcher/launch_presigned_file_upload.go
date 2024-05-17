package launcher

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/hse-experiments-platform/launcher/internal/pkg/domain"
	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateS3Objectname(filename string) string {
	return "usercontent_" + filename + "_" + randStringBytes(5)
}

func (s *launcherService) LaunchFilePresignedUpload(ctx context.Context, req *pb.LaunchFilePresignedUploadRequest) (*pb.LaunchFilePresignedUploadResponse, error) {
	if req.GetFilename() == "" {
		return nil, status.Error(codes.InvalidArgument, "filename must be not empty")
	}

	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("input json.Marshal: %w", err)
	}
	launchID, err := s.commonDB.CreateLaunch(ctx, db.CreateLaunchParams{
		LaunchType:   pb.LaunchType_LaunchTypeFileUpload.String(),
		UserID:       userID,
		Name:         req.GetLaunchInfo().GetName(),
		Description:  req.GetLaunchInfo().GetDescription(),
		LaunchStatus: pb.LaunchStatus_LaunchStatusNotStarted.String(),
		Input:        bytes,
	})
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.CreateLaunch: %w", err)
	}

	u, err := s.s3Storage.GetObjectUploadLink(ctx, domain.GetBucketName(userID), generateS3Objectname(req.GetFilename()))
	if err != nil {
		return nil, fmt.Errorf("s.s3Storage.GetObjectUploadLink: %w", err)
	}

	if err = s.launcher.Launch(launchID, s.fileUploadLauncher(u)); err != nil {
		return nil, fmt.Errorf("s.launcher.Launch: %w", err)
	}

	return &pb.LaunchFilePresignedUploadResponse{
		LaunchID:  launchID,
		UploadURL: u.String(),
	}, nil
}

func (s *launcherService) fileUploadLauncher(u *url.URL) func(context.Context, int64) ([]byte, error) {
	return func(ctx context.Context, launchID int64) ([]byte, error) {
		if bytes, errM := json.Marshal(u); errM != nil {
			return nil, errM
		} else {
			return bytes, nil
		}
	}
}

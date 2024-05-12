package launcher

import (
	bytes2 "bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hse-experiments-platform/launcher/internal/pkg/domain"
	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *launcherService) LaunchDatasetUpload(ctx context.Context, req *pb.LaunchDatasetUploadRequest) (*pb.LaunchDatasetUploadResponse, error) {
	if req.GetDatasetID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "datasetID must be not 0")
	}

	userID, err := s.checkDatasetAccess(ctx, req.GetDatasetID())
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("input json.Marshal: %w", err)
	}

	launchID, err := s.commonDB.CreateLaunch(ctx, db.CreateLaunchParams{
		LaunchType:   pb.LaunchType_LaunchTypeDatasetUpload.String(),
		UserID:       userID,
		Name:         req.GetLaunchInfo().GetName(),
		Description:  req.GetLaunchInfo().GetDescription(),
		LaunchStatus: pb.LaunchStatus_LaunchStatusNotStarted.String(),
		Input:        bytes,
	})
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.CreateLaunch: %w", err)
	}

	if err = s.launcher.Launch(launchID, s.datasetUploadLauncher(userID, req)); err != nil {
		return nil, fmt.Errorf("s.launcher.Launch: %w", err)
	}

	return &pb.LaunchDatasetUploadResponse{
		LaunchID: launchID,
	}, nil
}

func (s *launcherService) datasetUploadLauncher(userID int64, req *pb.LaunchDatasetUploadRequest) func(context.Context, int64) ([]byte, error) {
	return func(ctx context.Context, launchID int64) ([]byte, error) {
		body, err := makeUploadRequest(launchID, userID, req)
		if err != nil {
			return nil, fmt.Errorf("makeUploadRequest: %w", err)
		}

		var bytes []byte
		var errM error

		resp, err := s.client.SendDatasetUploadTask(ctx, body)
		if resp != nil {
			if bytes, errM = json.Marshal(resp); errM != nil {
				log.Error().Err(errM).Msg("json.Marshal")
			}
		}
		if err != nil {
			return bytes, fmt.Errorf("s.client.SendDatasetSetTypesTask: %w", err)
		}

		log.Debug().Str("resp", fmt.Sprint(resp)).Msg("response")

		return bytes, nil
	}
}

func makeUploadRequest(launchID, userID int64, req *pb.LaunchDatasetUploadRequest) (io.Reader, error) {
	var body struct {
		LaunchID  string `json:"launch_id"`
		UserID    string `json:"user_id"`
		DatasetID string `json:"dataset_id"`
		URL       string `json:"url"`
	}

	body.LaunchID = fmt.Sprintf("%d", launchID)
	body.UserID = domain.GetBucketName(userID)
	body.DatasetID = domain.GetObjectName(req.GetDatasetID())
	body.URL = req.GetUrl()

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	log.Debug().Str("body", string(bytes)).Msg("request")

	return bytes2.NewReader(bytes), nil
}

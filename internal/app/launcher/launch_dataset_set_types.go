package launcher

import (
	bytes2 "bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/hse-experiments-platform/launcher/internal/pkg/domain"
	"github.com/hse-experiments-platform/launcher/internal/pkg/storage/db"
	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *launcherService) LaunchDatasetSetTypes(ctx context.Context, req *pb.LaunchDatasetSetTypesRequest) (*pb.LaunchDatasetSetTypesResponse, error) {
	if req.GetDatasetID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "datasetID must be not 0")
	}

	userID, err := s.checkDatasetAccess(ctx, req.GetDatasetID())
	if err != nil {
		return nil, err
	}

	launchID, err := s.commonDB.CreateLaunch(ctx, db.CreateLaunchParams{
		LaunchType:   pb.LaunchType_LaunchTypeDatasetSetTypes.String(),
		UserID:       userID,
		Name:         req.GetLaunchInfo().GetName(),
		Description:  req.GetLaunchInfo().GetDescription(),
		LaunchStatus: pb.LaunchStatus_LaunchStatusNotStarted.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.CreateLaunch: %w", err)
	}

	if err = s.launcher.Launch(launchID, s.datasetSetTypesLauncher(userID, req)); err != nil {
		return nil, fmt.Errorf("s.launcher.Launch: %w", err)
	}

	return &pb.LaunchDatasetSetTypesResponse{
		LaunchID: launchID,
	}, nil
}

func (s *launcherService) datasetSetTypesLauncher(userID int64, req *pb.LaunchDatasetSetTypesRequest) func(context.Context, int64) error {
	return func(ctx context.Context, launchID int64) error {
		body, err := makeSetTypesRequest(launchID, userID, req)
		if err != nil {
			return fmt.Errorf("makeSetTypesRequest: %w", err)
		}

		err = s.client.SendDatasetSetTypesTask(ctx, body)
		if err != nil {
			return fmt.Errorf("s.client.SendDatasetSetTypesTask: %w", err)
		}

		return nil
	}
}

func makeSetTypesRequest(launchID, userID int64, req *pb.LaunchDatasetSetTypesRequest) (io.Reader, error) {
	type columnSettings struct {
		ColumnType        string `json:"column_type"`
		FillingTechnique  string `json:"filling_technique"`
		AggregateFunction string `json:"aggregate_function,omitempty"`
		FillingValue      string `json:"filling_value,omitempty"`
	}

	var body struct {
		LaunchID  string                    `json:"launch_id"`
		UserID    string                    `json:"user_id"`
		DatasetID string                    `json:"dataset_id"`
		Schema    map[string]columnSettings `json:"schema"`
	}

	body.LaunchID = fmt.Sprint(launchID)
	body.UserID = domain.GetBucketName(userID)
	body.DatasetID = domain.GetObjectName(req.GetDatasetID())
	body.Schema = make(map[string]columnSettings, len(req.GetColumnTypes()))
	for columnName, col := range req.GetColumnTypes() {
		colType, _ := strings.CutPrefix("ColumnType", col.GetType().String())
		fillingTechnique, _ := strings.CutPrefix("FillingTechnique", col.GetFillingTechnique().String())
		aggregateFunction, _ := strings.CutPrefix("AggregateFunction", col.GetAggregateFunction().String())

		body.Schema[columnName] = columnSettings{
			ColumnType:        colType,
			FillingTechnique:  fillingTechnique,
			AggregateFunction: aggregateFunction,
			FillingValue:      col.GetFillingValue(),
		}
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return bytes2.NewReader(bytes), nil
}

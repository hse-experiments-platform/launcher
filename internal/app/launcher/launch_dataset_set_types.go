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

func (s *launcherService) LaunchDatasetSetTypes(ctx context.Context, req *pb.LaunchDatasetSetTypesRequest) (*pb.LaunchDatasetSetTypesResponse, error) {
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
		LaunchType:   pb.LaunchType_LaunchTypeDatasetSetTypes.String(),
		UserID:       userID,
		Name:         req.GetLaunchInfo().GetName(),
		Description:  req.GetLaunchInfo().GetDescription(),
		LaunchStatus: pb.LaunchStatus_LaunchStatusNotStarted.String(),
		Input:        bytes,
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

func (s *launcherService) datasetSetTypesLauncher(userID int64, req *pb.LaunchDatasetSetTypesRequest) func(context.Context, int64) ([]byte, error) {
	return func(ctx context.Context, launchID int64) ([]byte, error) {
		body, err := makeSetTypesRequest(launchID, userID, req)
		if err != nil {
			return nil, fmt.Errorf("makeSetTypesRequest: %w", err)
		}

		err = s.client.SendDatasetSetTypesTask(ctx, body)
		if err != nil {
			return nil, fmt.Errorf("s.client.SendDatasetSetTypesTask: %w", err)
		}

		return []byte("{}"), nil
	}
}

func makeSetTypesRequest(launchID, userID int64, req *pb.LaunchDatasetSetTypesRequest) (io.Reader, error) {
	type emptiesSettings struct {
		Technique         string `json:"technique"`
		ConstantValue     string `json:"constant_value,omitempty"`
		AggregateFunction string `json:"aggregate_function,omitempty"`
	}

	type columnSettings struct {
		ColumnType      string          `json:"column_type"`
		EmptiesSettings emptiesSettings `json:"empties_settings"`
	}

	var body struct {
		LaunchID     string                    `json:"launch_id"`
		UserID       string                    `json:"user_id"`
		DatasetID    string                    `json:"dataset_id"`
		NewDatasetID string                    `json:"new_dataset_id"`
		Settings     map[string]columnSettings `json:"settings"`
	}

	body.LaunchID = fmt.Sprint(launchID)
	body.UserID = domain.GetBucketName(userID)
	body.DatasetID = domain.GetObjectName(req.GetDatasetID())
	body.NewDatasetID = body.DatasetID
	body.Settings = make(map[string]columnSettings, len(req.GetColumnTypes()))
	for columnName, col := range req.GetColumnTypes() {
		colType := convertColumnType(col.GetType())
		fillingTechnique := convertFillingTechnique(col.GetFillingTechnique())
		aggregateFunction := convertAggregateFunction(col.GetAggregateFunction())

		body.Settings[columnName] = columnSettings{
			ColumnType: colType,
			EmptiesSettings: emptiesSettings{
				Technique:         fillingTechnique,
				ConstantValue:     col.GetFillingValue(),
				AggregateFunction: aggregateFunction,
			},
		}
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	log.Debug().Str("set_types_request", string(bytes)).Msg("set types request")

	return bytes2.NewReader(bytes), nil
}

func convertColumnType(columnType pb.ColumnType) string {
	switch columnType {
	case pb.ColumnType_ColumnTypeInteger:
		return "Int"
	case pb.ColumnType_ColumnTypeFloat:
		return "Float"
	case pb.ColumnType_ColumnTypeCategorial:
		return "Categorial"
	case pb.ColumnType_ColumnTypeDropped:
		return "Dropped"
	default:
		return ""
	}
}

func convertFillingTechnique(fillingTechnique pb.FillingTechnique) string {
	switch fillingTechnique {
	case pb.FillingTechnique_FillingTechniqueConstant:
		return "Constant"
	case pb.FillingTechnique_FillingTechniqueTypeDefault:
		return "TypeDefault"
	case pb.FillingTechnique_FillingTechniqueAggregateFunction:
		return "AggregateFunction"
	case pb.FillingTechnique_FillingTechniqueDeleteRow:
		return "DeleteRow"
	default:
		return ""
	}
}

func convertAggregateFunction(aggregateFunction pb.AggregateFunction) string {
	switch aggregateFunction {
	case pb.AggregateFunction_AggregateFunctionMean:
		return "Mean"
	case pb.AggregateFunction_AggregateFunctionMedian:
		return "Median"
	case pb.AggregateFunction_AggregateFunctionMin:
		return "Min"
	case pb.AggregateFunction_AggregateFunctionMax:
		return "Max"
	case pb.AggregateFunction_AggregateFunctionMostFrequent:
		return "MostFrequent"
	default:
		return ""
	}
}

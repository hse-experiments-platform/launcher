package launcher

import (
	"context"
	"fmt"

	pb "github.com/hse-experiments-platform/launcher/pkg/launcher"
)

func (s *launcherService) GetTrainResults(ctx context.Context, req *pb.GetTrainResultsRequest) (*pb.GetTrainResultsResponse, error) {
	resp := &pb.GetTrainResultsResponse{}

	metrics, err := s.commonDB.GetTrainedModelMetrics(ctx, int64(req.GetLaunchID()))
	if err != nil {
		return nil, fmt.Errorf("s.commonDB.GetTrainedModelMetrics: %w", err)
	}

	for _, metric := range metrics {
		resp.Metrics = append(resp.Metrics, &pb.Metric{
			Id:          uint64(metric.MetricID),
			Name:        metric.Name,
			Description: metric.Description,
			Value:       string(metric.Value),
		})
	}

	return resp, nil
}

package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	datasetSetTypesPath = "/set_column_types"
	datasetUploadPath   = "/upload"
	trainPath           = "/train"
)

type WorkersClient interface {
	SendDatasetSetTypesTask(ctx context.Context, body io.Reader) error
	SendDatasetUploadTask(ctx context.Context, body io.Reader) ([]string, error)
	SendTrainTask(ctx context.Context, body io.Reader) (string, error)
}

type workersClient struct {
	baseUrl string

	c http.Client
}

func NewWorkersClient(baseUrl string) WorkersClient {
	return &workersClient{
		c:       http.Client{},
		baseUrl: baseUrl,
	}
}

func (w *workersClient) do(ctx context.Context, body io.Reader, path string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", w.baseUrl+path, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	resp, err := w.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("w.c.Do: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return resp.Body, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func (w *workersClient) SendDatasetSetTypesTask(ctx context.Context, body io.Reader) error {
	resp, err := w.do(ctx, body, datasetSetTypesPath)
	if err != nil {
		if resp != nil {
			if bytes, err2 := io.ReadAll(resp); err2 != nil {
				log.Error().Err(err2).Msg("cannot read response body")
			} else {
				log.Error().Err(err).Str("resp", string(bytes)).Msg("error response body")
			}
		}

		return fmt.Errorf("w.do: %w", err)
	}
	defer func() { _ = resp.Close() }()

	return nil
}

func (w *workersClient) SendDatasetUploadTask(ctx context.Context, body io.Reader) ([]string, error) {
	resp, err := w.do(ctx, body, datasetUploadPath)
	if err != nil {
		if resp != nil {
			if bytes, err2 := io.ReadAll(resp); err2 != nil {
				log.Error().Err(err2).Msg("cannot read response body")
			} else {
				log.Error().Err(err).Str("resp", string(bytes)).Msg("error response body")
			}
		}

		return nil, fmt.Errorf("w.do: %w", err)
	}
	defer func() { _ = resp.Close() }()

	bytes, err := io.ReadAll(resp)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	var output []string
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}

func (w *workersClient) SendTrainTask(ctx context.Context, body io.Reader) (string, error) {
	resp, err := w.do(ctx, body, trainPath)
	if err != nil {
		if resp != nil {
			if bytes, err2 := io.ReadAll(resp); err2 != nil {
				log.Error().Err(err2).Msg("cannot read response body")
			} else {
				log.Error().Err(err).Str("resp", string(bytes)).Msg("error response body")
			}
		}

		return "", fmt.Errorf("w.do: %w", err)
	}
	defer func() { _ = resp.Close() }()

	bytes, err := io.ReadAll(resp)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll: %w", err)
	}

	var output string
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal: %w", err)
	}

	return output, nil
}

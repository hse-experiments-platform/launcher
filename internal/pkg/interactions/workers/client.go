package workers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	datasetSetTypesPath = "/set_column_types"
	datasetUploadPath   = "/upload"
)

type WorkersClient interface {
	SendDatasetSetTypesTask(ctx context.Context, body io.Reader) error
	SendDatasetUploadTask(ctx context.Context, body io.Reader) ([]byte, error)
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
	defer func() { _ = resp.Close() }()
	if err != nil {
		if bytes, err := io.ReadAll(resp); err != nil {
			log.Error().Err(err).Msg("cannot read response body")
		} else {
			log.Error().Err(err).Str("resp", string(bytes)).Msg("error response body")
		}

		return fmt.Errorf("w.do: %w", err)
	}

	return nil
}

func (w *workersClient) SendDatasetUploadTask(ctx context.Context, body io.Reader) ([]byte, error) {
	resp, err := w.do(ctx, body, datasetUploadPath)
	defer func() { _ = resp.Close() }()
	if err != nil {
		if bytes, err2 := io.ReadAll(resp); err2 != nil {
			log.Error().Err(err2).Msg("cannot read response body")
			return nil, err2
		} else {
			log.Error().Err(err).Str("resp", string(bytes)).Msg("error response body")
			return bytes, fmt.Errorf("w.do: %w", err)
		}

	}

	return io.ReadAll(resp)
}

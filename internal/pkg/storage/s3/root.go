package s3

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinioStorage struct {
	client *minio.Client
}

func NewMinioStorage(client *minio.Client) *MinioStorage {
	return &MinioStorage{client: client}
}

func (s *MinioStorage) CheckOrCreateBucket(ctx context.Context, bucketName string) error {
	exists, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("s.client.BucketExists: %w", err)
	}

	if !exists {
		err = s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("s.client.MakeBucket: %w", err)
		}
	}

	return nil
}

func (s *MinioStorage) GetObjectDownloadLink(ctx context.Context, bucketName, objectName string) (*url.URL, error) {
	return s.client.PresignedGetObject(ctx, bucketName, objectName, time.Hour*24*7, nil)
}

package dai

import (
	"context"
)

type FileStore interface {
	UploadFile(ctx context.Context, file []byte, key string) (string, error)
	GetPresignedObjectURL(ctx context.Context, key string) (string, error)
	DeleteFile(ctx context.Context, fileName string) error
	ParallelGetPresignedObjectURL(ctx context.Context, input []ParallelGetPresignedObjectURLInput) (map[string]string, error)
}

type ParallelGetPresignedObjectURLInput struct {
	HackathonID string
	Key         string
}

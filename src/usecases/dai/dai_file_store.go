package dai

import (
	"context"
)

type FileStore interface {
	UploadFile(ctx context.Context, file []byte, key string) (string, error)
	DeleteFile(ctx context.Context, fileName string) error
}

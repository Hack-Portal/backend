package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/gateways/repository/datasource"
)

type EtcUsecase interface {
	GetFramework(ctx context.Context, limit int32) ([]repository.Framework, error)
	GetLocat(ctx context.Context) ([]repository.Locate, error)
	GetTechTag(ctx context.Context) ([]repository.TechTag, error)
}

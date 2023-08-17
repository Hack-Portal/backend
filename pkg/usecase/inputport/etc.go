package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type EtcUsecase interface {
	GetFramework(ctx context.Context) ([]repository.Framework, error)
	GetLocat(ctx context.Context) ([]repository.Locate, error)
	GetTechTag(ctx context.Context) ([]repository.TechTag, error)
	GetStatusTag(ctx context.Context) ([]repository.StatusTag, error)
}

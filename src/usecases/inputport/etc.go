package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

type EtcUsecase interface {
	GetFramework(ctx context.Context) ([]repository.Framework, error)
	GetLocat(ctx context.Context) ([]repository.Locate, error)
	GetTechTag(ctx context.Context) ([]repository.TechTag, error)
	GetStatusTag(ctx context.Context) ([]repository.StatusTag, error)
	ListRoles(ctx context.Context) ([]repository.Role, error)
}

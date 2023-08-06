package domain

import (
	"context"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

type EtcUsecase interface {
	GetFramework(ctx context.Context, limit int32) ([]db.Frameworks, error)
	GetLocat(ctx context.Context) ([]db.Locates, error)
	GetTechTag(ctx context.Context) ([]db.TechTags, error)
}

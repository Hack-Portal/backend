package domain

import (
	"context"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

type CreateRateRequestBody struct {
	Rate int32 `json:"rate"`
}

type ListRateParams struct {
	PageSize int32 `form:"page_size"`
	PageId   int32 `form:"page_id"`
}

type RateUsecase interface {
	CreateRateEntry(ctx context.Context, body db.CreateRateParams) (db.RateEntries, error)
	ListRateEntry(ctx context.Context, id string, query ListRateParams) ([]db.RateEntries, error)
}

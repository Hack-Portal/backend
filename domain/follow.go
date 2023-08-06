package domain

import (
	"context"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

type CreateFollowRequestBody struct {
	ToUserID string `json:"to_user_id" binding:"required"`
}

type RemoveFollowRequestQueries struct {
	ToUserID   string `json:"to_user_id" binding:"required"`
	FromUserID string `json:"from_user_id" binding:"required"`
}

type GetFollowRequestQueries struct {
	Mode     bool   `form:"mode"`
	PageSize string `form:"page_size"`
	PageID   string `form:"page_id"`
}

type FollowResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}

type FollowUsecase interface {
	CreateFollow(ctx context.Context, body db.CreateFollowParams) (result FollowResponse, err error)
	RemoveFollow(ctx context.Context, body db.SoftRemoveBookmarkParams) (err error)
	GetFollowByToID(ctx context.Context, ID string) (result []FollowResponse, err error)
}

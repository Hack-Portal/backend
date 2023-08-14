package domain

import (
	"time"
)

type CreateBookmarkRequest struct {
	AccountID string `json:"account_id"`
	Opus      int32  `json:"opus"`
}

type BookmarkResponse struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
}

type BookmarkRequestWildCard struct {
	AccountID string `uri:"account_id"`
}
type RemoveBookmarkRequestQueries struct {
	Opus int32 `query:"opus" binding:"required"`
}
type ListBookmarkRequestQueries struct {
	PageSize int32 `form:"page_size" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required"`
}

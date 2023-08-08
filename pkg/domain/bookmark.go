package domain

import (
	"time"
)

type CreateBookmarkRequest struct {
	UserID      string `json:"user_id"`
	HackathonID int32  `json:"hackathon_id"`
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
	UserID string `uri:"user_id"`
}
type RemoveBookmarkRequestQueries struct {
	HackathonID int32 `query:"hackathon_id" binding:"required"`
}
type ListBookmarkRequestQueries struct {
	PageSize int32 `form:"page_size" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required"`
}

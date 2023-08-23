package domain

import (
	"time"
)

type CreateLikeRequest struct {
	AccountID string `json:"account_id"`
	Opus      int32  `json:"opus"`
}

type LikeResponse struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
}

type LikeRequestWildCard struct {
	AccountID string `uri:"account_id"`
}
type RemoveLikeRequestQueries struct {
	Opus int32 `form:"opus" binding:"required"`
}

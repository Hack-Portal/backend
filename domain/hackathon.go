package domain

import (
	"context"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

type HackathonRequestWildCard struct {
	HackathonID int32 `uri:"hackathon_id"`
}

type CreateHackathonRequestBody struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
	StatusTags  []int32   `json:"status_tags"`
}

type CreateHackathonParams struct {
	CreateHackathonRequestBody
	Image      []byte  `json:"image"`
	StatusTags []int32 `json:"status_tags"`
}

type HackathonResponses struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	StatusTags []db.StatusTags `json:"status_tags"`
}

type ListHackathonsParams struct {
	PageSize int32 `form:"page_size"`
	PageId   int32 `form:"page_id"`
	Expired  bool  `form:"expired"`
}

type ListHackathonsResponses struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	StatusTags []db.StatusTags `json:"status_tags"`
}

type HackathonUsecase interface {
	CreateHackathon(ctx context.Context, body CreateHackathonParams) (result HackathonResponses, err error)
	GetHackathon(ctx context.Context, id int32) (result HackathonResponses, err error)
	ListHackathons(ctx context.Context, query ListHackathonsParams) (result []ListHackathonsResponses, err error)
}

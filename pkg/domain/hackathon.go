package domain

import (
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type HackathonRequestWildCard struct {
	HackathonID int32 `uri:"hackathon_id"`
}

type CreateHackathonRequestBody struct {
	Name        string    `form:"name"`
	Description string    `form:"description"`
	Link        string    `form:"link"`
	Expired     time.Time `form:"expired"`
	StartDate   time.Time `form:"start_date"`
	Term        int32     `form:"term"`
	StatusTags  []int32   `form:"status_tags"`
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

	StatusTags []repository.StatusTag `json:"status_tags"`
}

type ListHackathonsParams struct {
	ListRequest
	Expired bool `form:"expired"`
}

type ListHackathonsResponses struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	StatusTags []repository.StatusTag `json:"status_tags"`
}

type CreateHackathonParams struct {
	Hackathon  repository.CreateHackathonsParams
	StatusTags []int32 `json:"status_tags"`
}

package entities

import (
	"time"
)

type Hackathon struct {
	HackathonID string    `json:"hackathon_id" gorm:"primaryKey" `
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int       `json:"term"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDelete  bool      `json:"is_delete"`
}

type HackathonStatusTag struct {
	HackathonID string `json:"hackathon_id"`
	StatusID    int    `json:"status_id"`
}

type StatusTag struct {
	StatusID int    `json:"status_id"`
	Status   string `json:"status"`
}

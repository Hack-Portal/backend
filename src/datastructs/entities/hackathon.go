package entities

import "time"

// ここでは、データベースのテーブルを定義する

type Hackathon struct {
	HackathonID string    `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsDelete    bool      `json:"is_delete"`
}

type HackathonStatusTag struct {
	HackathonID string `json:"hackathon_id"`
	StatusID    int32  `json:"status_id"`
}

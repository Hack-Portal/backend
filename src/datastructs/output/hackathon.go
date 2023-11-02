package output

import (
	"temp/src/datastructs/entities"
	"time"
)

// ここでは、返す構造体を定義する
// 所謂レスポンスボディ等の構造体を定義する

type HackathonCreate struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	StatusTags []entities.StatusTag `json:"status_tags"`
}

type HackathonRead struct{}
type HackathonUpdate struct{}
type HackathonDelete struct{}

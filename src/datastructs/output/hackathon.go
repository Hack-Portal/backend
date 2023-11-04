package output

import (
	"time"

	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
)

// ここでは、返す構造体を定義する
// 所謂レスポンスボディ等の構造体を定義する

type CreateHackathon struct {
	Message string
}

type ReadAllHackathon struct {
	HackathonID string               `json:"hackathon_id"`
	Name        string               `form:"name"`
	Link        string               `form:"link"`
	Expired     time.Time            `form:"expired"`
	StartDate   time.Time            `form:"start_date"`
	Term        int32                `form:"term"`
	StatusTags  []entities.StatusTag `form:"status_tags"`
}

type HackathonUpdate struct{}
type HackathonDelete struct{}

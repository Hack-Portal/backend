package request

import "time"

type CreateHackathon struct {
	Name      string    `form:"name" validate:"required"`
	Link      string    `form:"link" validate:"required"`
	Expired   time.Time `form:"expired" validate:"required"`
	StartDate time.Time `form:"start_date" validate:"required"`
	Term      int       `form:"term" validate:"required"`
	Statuses  []int64   `form:"statuses[]"`
}

type GetHackathon struct {
	HackathonID string `param:"hackathon_id" validate:"required"`
}

type ListHackathon struct {
	PageSize int `query:"page_size"`
	PageID   int `query:"page_id"`

	// タグ
	Tags []int64 `query:"tags"`
	// 新着かどうか？
	New bool `query:"new"`
	// 期間が長いかどうか？
	LongTerm bool `query:"long_term"`
	// 締め切りが近いかどうか？
	NearDeadline bool `query:"near_deadline"`
}

type DeleteHackathon struct {
	HackathonID string `param:"hackathon_id" validate:"required"`
}

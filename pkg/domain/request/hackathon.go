package request

import "time"

type HackathonWildCard struct {
	HackathonID int32 `uri:"hackathon_id"`
}

type ListHackathons struct {
	ListRequest
	Expired bool `form:"expired"`
}

type CreateHackathon struct {
	Name        string    `form:"name"`
	Description string    `form:"description"`
	Link        string    `form:"link"`
	Expired     time.Time `form:"expired"`
	StartDate   time.Time `form:"start_date"`
	Term        int32     `form:"term"`
	StatusTags  []int32   `form:"status_tags"`
}

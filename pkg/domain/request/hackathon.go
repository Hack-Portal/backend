package request

import "time"

type HackathonRequestWildCard struct {
	HackathonID int32 `uri:"hackathon_id"`
}

type ListHackathonsRequest struct {
	ListRequest
	Expired bool `form:"expired"`
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

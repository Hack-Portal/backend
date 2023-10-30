package request

import "time"

type CreateHackathon struct {
	Name        string    `form:"name"`
	Description string    `form:"description"`
	Link        string    `form:"link"`
	Expired     time.Time `form:"expired"`
	StartDate   time.Time `form:"start_date"`
	Term        int       `form:"term"`
	StatusTags  string    `form:"status_tags"`
}

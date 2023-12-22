package request

import "time"

type CreateHackathon struct {
	Name      string    `json:"name"`
	Link      string    `json:"link"`
	Expired   time.Time `json:"expired"`
	StartDate time.Time `json:"start_date"`
	Term      int       `json:"term"`
}

package request

import "time"

type CreateHackathon struct {
	Name      string    `json:"name" validate:"required"`
	Link      string    `json:"link" validate:"required"`
	Expired   time.Time `json:"expired" validate:"required"`
	StartDate time.Time `json:"start_date" validate:"required"`
	Term      int       `json:"term" validate:"required"`
}

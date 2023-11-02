package entities

import "time"

type ProposalHackathon struct {
	ProposalID string    `json:"proposal_id"`
	Name       string    `json:"name"`
	Icon       string    `json:"icon"`
	Link       string    `json:"link"`
	Expired    time.Time `json:"expired"`
	StartDate  time.Time `json:"start_date"`
	Term       int32     `json:"term"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProposalHackathonStatusTag struct {
	ProposalID string `json:"proposal_id"`
	StatusID   int32  `json:"status_id"`
}

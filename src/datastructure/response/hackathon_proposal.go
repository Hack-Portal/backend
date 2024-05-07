package response

import "time"

type CreateHackathonProposal struct {
	HackathonProposalID int    `json:"hackathon_proposal_id"`
	URL                 string `json:"url"`
}

type GetHackathonProposal struct {
	HackathonProposalID int       `json:"hackathon_proposal_id"`
	URL                 string    `json:"url"`
	IsApproved          bool      `json:"is_approved"`
	CreatedAt           time.Time `json:"created_at"`
}

type UpdateHackathonProposal struct {
	HackathonProposalID int       `json:"hackathon_proposal_id"`
	URL                 string    `json:"url"`
	IsApproved          bool      `json:"is_approved"`
	CreatedAt           time.Time `json:"created_at"`
}

type DeleteHackathonProposal struct {
	HackathonProposalID int `json:"hackathon_proposal_id"`
}

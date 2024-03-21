package response

import "time"

type CreateHackathonProposal struct {
	HackathonProposalID int       `json:"hackathon_proposal_id"`
	URL                 string    `json:"url"`
	IsApproved          bool      `json:"is_approved"`
	CreatedAt           time.Time `json:"created_at"`
}

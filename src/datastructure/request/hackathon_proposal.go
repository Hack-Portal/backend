package request

type CreateHackathonProposal struct {
	URL string `json:"url"`
}

type ListHackathonProposal struct {
	PageSize int `query:"page_size"`
	PageID   int `query:"page_id"`
}

type GetHackathonProposal struct {
	HackathonProposalID int `json:"hackathon_proposal_id"`
}

type UpdateHackathonProposal struct {
	HackathonProposalID int    `json:"hackathon_proposal_id"`
	URL                 string `json:"url"`
	IsApproved          bool   `json:"isApproved"`
}

type DeleteHackathonProposal struct {
	HackathonProposalID int `json:"hackathon_proposal_id"`
}

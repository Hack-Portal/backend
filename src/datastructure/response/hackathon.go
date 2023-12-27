package response

type CreateHackathon struct {
	HackathonID string `json:"hackathon_id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Link        string `json:"link"`
	Expired     string `json:"expired"`
	StartDate   string `json:"start_date"`
	Term        int    `json:"term"`

	StatusTags []*StatusTag `json:"status_tags"`
}

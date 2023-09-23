package request

type AccountRequestWildCard struct {
	AccountID string `uri:"account_id"`
}

type CreateAccountRequest struct {
	AccountID       string `form:"account_id" binding:"required"`
	Username        string `form:"username" binding:"required"`
	ExplanatoryText string `form:"explanatory_text"`
	LocateID        int32  `form:"locate_id" binding:"required"`
	ShowLocate      bool   `form:"show_locate"`
	ShowRate        bool   `form:"show_rate" `

	TechTags   string `form:"tech_tags"`
	Frameworks string `form:"frameworks"`
}

type UpdateAccountRequest struct {
	Username        string `form:"username"`
	ExplanatoryText string `form:"explanatory_text"`
	LocateID        int32  `form:"locate_id"`
	ShowLocate      bool   `form:"show_locate"`
	ShowRate        bool   `form:"show_rate"`
	GithubLink      string `form:"github_link"`
	TwitterLink     string `form:"twitter_link"`
	DiscordLink     string `form:"discord_link"`

	TechTags   string `form:"tech_tags"`
	Frameworks string `form:"frameworks"`
}

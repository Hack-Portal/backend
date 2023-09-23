package params

type CreatePastWork struct {
	Name               string   `json:"name"`
	ThumbnailImage     string   `json:"thumbnail_image"`
	ExplanatoryText    string   `json:"explanatory_text"`
	PastWorkTags       []int32  `json:"past_work_tags"`
	PastWorkFrameworks []int32  `json:"past_work_frameworks"`
	AccountPastWorks   []string `json:"account_past_works"`
}

type UpdatePastWork struct {
	Opus               int32    `form:"opus"`
	Name               string   `form:"name"`
	ExplanatoryText    string   `form:"explanatory_text"`
	PastWorkTags       []int32  `form:"past_work_tags"`
	PastWorkFrameworks []int32  `form:"past_work_frameworks"`
	AccountPastWorks   []string `form:"account_past_works"`
}

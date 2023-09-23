package request

type ListRequest struct {
	PageSize int32 `form:"page_size"`
	PageID   int32 `form:"page_id"`
}
type PastWorksRequestWildCard struct {
	Opus int32 `uri:"opus"`
}

type PastWorkRequestBody struct {
	Name               string `form:"name"`
	ExplanatoryText    string `form:"explanatory_text"`
	PastWorkTags       string `form:"past_work_tags"`
	PastWorkFrameworks string `form:"past_work_frameworks"`
	AccountPastWorks   string `form:"account_past_works"`
}

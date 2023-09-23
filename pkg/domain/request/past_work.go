package request

type PastWorksWildCard struct {
	Opus int32 `uri:"opus"`
}

type PastWork struct {
	Name               string `form:"name"`
	ExplanatoryText    string `form:"explanatory_text"`
	PastWorkTags       string `form:"past_work_tags"`
	PastWorkFrameworks string `form:"past_work_frameworks"`
	AccountPastWorks   string `form:"account_past_works"`
}

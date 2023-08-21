package domain

import (
	"database/sql"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type PastWorksRequestWildCard struct {
	Opus int32 `uri:"opus"`
}

type CreatePastWorkRequestBody struct {
	Name               string   `form:"name"`
	ExplanatoryText    string   `form:"explanatory_text"`
	PastWorkTags       []int32  `form:"past_work_tags"`
	PastWorkFrameworks []int32  `form:"past_work_frameworks"`
	AccountPastWorks   []string `form:"account_past_works"`
}

type CreatePastWorkResponse struct {
	Opus               int32                          `json:"opus"`
	Name               string                         `json:"name"`
	ThumbnailImage     []byte                         `json:"thumbnail_image"`
	ExplanatoryText    string                         `json:"explanatory_text"`
	PastWorkTags       []repository.PastWorkTag       `json:"past_work_tags"`
	PastWorkFrameworks []repository.PastWorkFramework `json:"past_work_frameworks"`
	AccountPastWorks   []repository.AccountPastWork   `json:"account_past_works"`
}

type PastWorkMembers struct {
	AccountID string `json:"account_id"`
	Icon      string `json:"icon"`
	Name      string `json:"name"`
}

type CreatePastWorkParams struct {
	Name               string        `json:"name"`
	ThumbnailImage     string        `json:"thumbnail_image"`
	ExplanatoryText    string        `json:"explanatory_text"`
	AwardDataID        sql.NullInt32 `json:"award_data_id"`
	PastWorkTags       []int32       `json:"past_work_tags"`
	PastWorkFrameworks []int32       `json:"past_work_frameworks"`
	AccountPastWorks   []string      `json:"account_past_works"`
}

type PastWorkResponse struct {
	PastWork   repository.PastWork    `json:"past_work"`
	TechTags   []repository.TechTag   `json:"tech_tags"`
	Frameworks []repository.Framework `json:"frameworks"`
	Members    []PastWorkMembers      `json:"members"`
}

type ListPastWorkResponse struct {
	Opus            int32                  `json:"opus"`
	Name            string                 `json:"name"`
	ExplanatoryText string                 `json:"explanatory_text"`
	TechTags        []repository.TechTag   `json:"tech_tags"`
	Frameworks      []repository.Framework `json:"frameworks"`
	Members         []PastWorkMembers      `json:"members"`
}

type UpdatePastWorkRequestBody struct {
	Opus               int32    `form:"opus"`
	Name               string   `form:"name"`
	ExplanatoryText    string   `form:"explanatory_text"`
	PastWorkTags       []int32  `form:"past_work_tags"`
	PastWorkFrameworks []int32  `form:"past_work_frameworks"`
	AccountPastWorks   []string `form:"account_past_works"`
}

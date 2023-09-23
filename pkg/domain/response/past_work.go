package response

import (
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type CreatePastWork struct {
	Opus               int32                          `json:"opus"`
	Name               string                         `json:"name"`
	ThumbnailImage     []byte                         `json:"thumbnail_image"`
	ExplanatoryText    string                         `json:"explanatory_text"`
	PastWorkTags       []repository.PastWorkTag       `json:"past_work_tags"`
	PastWorkFrameworks []repository.PastWorkFramework `json:"past_work_frameworks"`
	AccountPastWorks   []repository.AccountPastWork   `json:"account_past_works"`
}

type PastWork struct {
	Opus            int32     `json:"opus"`
	Name            string    `json:"name"`
	ThumbnailImage  string    `json:"thumbnail_image"`
	ExplanatoryText string    `json:"explanatory_text"`
	AwardDataID     int32     `json:"award_data_id"`
	CreateAt        time.Time `json:"create_at"`
	UpdateAt        time.Time `json:"update_at"`
	IsDelete        bool      `json:"is_delete"`

	TechTags   []repository.TechTag   `json:"tech_tags"`
	Frameworks []repository.Framework `json:"frameworks"`
	Members    []PastWorkMembers      `json:"members"`
}
type ListPastWork struct {
	Opus            int32                  `json:"opus"`
	Name            string                 `json:"name"`
	ExplanatoryText string                 `json:"explanatory_text"`
	TechTags        []repository.TechTag   `json:"tech_tags"`
	Frameworks      []repository.Framework `json:"frameworks"`
	Members         []PastWorkMembers      `json:"members"`
}

type PastWorkMembers struct {
	AccountID string `json:"account_id"`
	Icon      string `json:"icon"`
	Name      string `json:"name"`
}

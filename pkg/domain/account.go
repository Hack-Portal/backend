package domain

import (
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

// アカウントのパス用のクエリ
type AccountRequestWildCard struct {
	AccountID string `uri:"account_id"`
}

// アカウント作成のリクエストパラメータ
type CreateAccountRequest struct {
	AccountID       string `form:"account_id" binding:"required"`
	Username        string `form:"username" binding:"required"`
	ExplanatoryText string `form:"explanatory_text"`
	LocateID        int32  `form:"locate_id" binding:"required"`
	ShowLocate      bool   `form:"show_locate"`
	ShowRate        bool   `form:"show_rate" `

	TechTags   []int32 `form:"tech_tags"`
	Frameworks []int32 `form:"frameworks"`
}

// アカウント取得のレスポンス
type AccountResponses struct {
	AccountID       string `json:"account_id"`
	Username        string `json:"username"`
	Icon            string `json:"icon"`
	ExplanatoryText string `json:"explanatory_text"`
	Rate            int32  `json:"rate"`
	Email           string `json:"email"`
	Locate          string `json:"locate"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`

	TechTags   []repository.TechTag   `json:"tech_tags"`
	Frameworks []repository.Framework `json:"frameworks"`

	CreatedAt time.Time `json:"created_at"`
}

// アカウント更新のリクエストパラメータ
type UpdateAccountRequest struct {
	Username        string `form:"username"`
	ExplanatoryText string `form:"explanatory_text"`
	LocateID        int32  `form:"locate_id"`
	ShowLocate      bool   `form:"show_locate"`
	ShowRate        bool   `form:"show_rate"`

	TechTags   []int32 `form:"tech_tags"`
	Frameworks []int32 `form:"frameworks"`
}

type CreateAccountParams struct {
	AccountInfo         repository.CreateAccountsParams
	AccountTechTag      []int32
	AccountFrameworkTag []int32
}

type UpdateAccountParam struct {
	AccountInfo         repository.Account
	AccountTechTag      []int32
	AccountFrameworkTag []int32
}

type AccountRateResponse struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
	Icon      string `json:"icon"`
	Rate      int32  `json:"rate"`
}

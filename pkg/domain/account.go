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
	AccountID       string `json:"account_id" binding:"required"`
	Username        string `json:"username" binding:"required"`
	ExplanatoryText string `json:"explanatory_text"`
	LocateID        int32  `json:"locate_id" binding:"required"`
	ShowLocate      bool   `json:"show_locate" binding:"required"`
	ShowRate        bool   `json:"show_rate" binding:"required"`

	TechTags   []int32 `json:"tech_tags"`
	Frameworks []int32 `json:"frameworks"`
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
	Username        string `json:"username"`
	ExplanatoryText string `json:"explanatory_text"`
	LocateID        int32  `json:"locate_id"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`

	TechTags   []int32 `json:"tech_tags"`
	Frameworks []int32 `json:"frameworks"`
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

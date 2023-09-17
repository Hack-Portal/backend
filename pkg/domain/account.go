package domain

import (
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

	TechTags   string `form:"tech_tags"`
	Frameworks string `form:"frameworks"`
}

type CreateAccount struct {
	ReqBody    CreateAccountRequest
	TechTags   []int32
	Frameworks []int32
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
	GithubLink      string `json:"github_link"`
	TwitterLink     string `json:"twitter_link"`
	DiscordLink     string `json:"discord_link"`

	ShowLocate bool `json:"show_locate"`
	ShowRate   bool `json:"show_rate"`

	TechTags   []repository.TechTag   `json:"tech_tags"`
	Frameworks []repository.Framework `json:"frameworks"`
}

// アカウント更新のリクエストパラメータ
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

type GetJoinRoomResponse struct {
	RoomID string `json:"room_id"`
	Title  string `json:"title"`
}

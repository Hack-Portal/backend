package domain

import (
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

// アカウントのパス用のクエリ
type AccountRequestWildCard struct {
	UserID string `uri:"user_id"`
}

// アカウント作成のリクエストパラメータ
type CreateAccountRequest struct {
	UserID          string  `json:"user_id" binding:"required"`
	Username        string  `json:"username" binding:"required"`
	Icon            string  `json:"icon"`
	ExplanatoryText string  `json:"explanatory_text"`
	LocateID        int32   `json:"locate_id" binding:"required"`
	Password        string  `json:"password"`
	ShowLocate      bool    `json:"show_locate" binding:"required"`
	ShowRate        bool    `json:"show_rate" binding:"required"`
	TechTags        []int32 `json:"tech_tags"`
	Frameworks      []int32 `json:"frameworks"`
}

// アカウント取得のレスポンス
type AccountResponses struct {
	UserID          string `json:"user_id"`
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
	Rate            int32  `json:"rate"`
	HashedPassword  string `json:"hashed_password"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`
}

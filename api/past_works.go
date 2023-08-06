package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

// 過去作品の登録のリクエストパラメータ
type CreatePastWorkRequestBody struct {
	Name               string   `json:"name"`
	ThumbnailImage     []byte   `json:"thumbnail_image"`
	ExplanatoryText    string   `json:"explanatory_text"`
	PastWorkTags       []int32  `json:"past_work_tags"`
	PastWorkFrameworks []int32  `json:"past_work_frameworks"`
	AccountPastWorks   []string `json:"account_past_works"`
}
type CreatePastWorkResponse struct {
	Opus               int32                   `json:"opus"`
	Name               string                  `json:"name"`
	ThumbnailImage     []byte                  `json:"thumbnail_image"`
	ExplanatoryText    string                  `json:"explanatory_text"`
	PastWorkTags       []db.PastWorkTags       `json:"past_work_tags"`
	PastWorkFrameworks []db.PastWorkFrameworks `json:"past_work_frameworks"`
	AccountPastWorks   []db.AccountPastWorks   `json:"account_past_works"`
}

// 過去作品を登録する
func (server *Server) CreatePastWork(ctx *gin.Context) {
	var req CreatePastWorkRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreatePastWorkTxParams{
		Name:               req.Name,
		ThumbnailImage:     req.ThumbnailImage,
		ExplanatoryText:    req.ExplanatoryText,
		PastWorkTags:       req.PastWorkTags,
		PastWorkFrameworks: req.PastWorkFrameworks,
		AccountPastWorks:   req.AccountPastWorks,
	}
	result, err := server.store.CreatePastWorkTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := CreatePastWorkResponse{
		Opus:               result.Opus,
		Name:               result.Name,
		ThumbnailImage:     result.ThumbnailImage,
		ExplanatoryText:    result.ExplanatoryText,
		PastWorkTags:       result.PastWorkTags,
		PastWorkFrameworks: result.PastWorkFrameworks,
		AccountPastWorks:   result.AccountPastWorks,
	}
	ctx.JSON(http.StatusOK, response)
}

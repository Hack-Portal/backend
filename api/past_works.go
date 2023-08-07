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
// CreatePastWork godoc
// @Summary Create pastWork
// @Description create pastWork
// @Tags past_works
// @Accept  json
// @Produce  json
// @Param past_work body CreatePastWorkRequestBody true "past work"
// @Success 200 {object} CreatePastWorkResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /past_works [post]
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

// 過去作品取得
type PastWorksRequestWildCard struct {
	Opus int32 `uri:"opus"`
}

// GetPastWork godoc
// @Summary Get pastWork
// @Description get pastWork
// @Tags past_works
// @Accept  json
// @Produce  json
// @Param opus path int true "PastWorks API wildcard"
// @Success 200 {object} CreatePastWorkResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /past_works/{opus} [get]
func (server *Server) GetPastWork(ctx *gin.Context) {
	var req PastWorksRequestWildCard
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	pastWork, err := server.store.GetPastWorksByOpus(ctx, req.Opus)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	pastWorkTags, err := server.store.GetPastWorkTagsByOpus(ctx, req.Opus)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	pastWorkFrameworks, err := server.store.GetPastWorkFrameworksByOpus(ctx, req.Opus)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	accountPastWorks, err := server.store.GetAccountPastWorksByOpus(ctx, req.Opus)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := CreatePastWorkResponse{
		Opus:               pastWork.Opus,
		Name:               pastWork.Name,
		ThumbnailImage:     pastWork.ThumbnailImage,
		ExplanatoryText:    pastWork.ExplanatoryText,
		PastWorkTags:       pastWorkTags,
		PastWorkFrameworks: pastWorkFrameworks,
		AccountPastWorks:   accountPastWorks,
	}
	ctx.JSON(http.StatusOK, response)
}

// ListPastWorks godoc
// @Summary List pastWorks
// @Description list pastWorks
// @Tags past_works
// @Produce  json
// @Param page_size path int32 true "page_size"
// @Param page_id path int32 true "page_id"
// @Success 200 {array} CreatePastWorkResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /past_works [get]
func (server *Server) ListPastWorks(ctx *gin.Context) {
	var request ListHackathonsParams
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	pastWorks, err := server.store.ListPastWorks(ctx, db.ListPastWorksParams{
		Limit:  request.PageSize,
		Offset: (request.PageId - 1) * request.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []CreatePastWorkResponse
	for _, pastWork := range pastWorks {
		pastWorkTags, err := server.store.GetPastWorkTagsByOpus(ctx, pastWork.Opus)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		pastWorkFrameworks, err := server.store.GetPastWorkFrameworksByOpus(ctx, pastWork.Opus)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accountPastWorks, err := server.store.GetAccountPastWorksByOpus(ctx, pastWork.Opus)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		response = append(response, CreatePastWorkResponse{
			Opus:               pastWork.Opus,
			Name:               pastWork.Name,
			ExplanatoryText:    pastWork.ExplanatoryText,
			PastWorkTags:       pastWorkTags,
			PastWorkFrameworks: pastWorkFrameworks,
			AccountPastWorks:   accountPastWorks,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

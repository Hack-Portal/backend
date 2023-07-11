package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

// ハッカソンを作る時のリクエストパラメータ
type CreateHackathonParams struct {
	Name        string    `json:"Name"`
	Icon        []byte    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
	StatusTags  []int32   `json:"status_tags"`
}

// ハッカソンに関するレスポンス
type HackathonResponses struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        []byte    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	StatusTags []db.StatusTags
}

// ハッカソン作成
func (server *Server) CreateHackathon(ctx *gin.Context) {
	var request CreateHackathonParams
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateHackathonTxParams{
		Hackathons: db.Hackathons{
			Name:        request.Name,
			Icon:        request.Icon,
			Description: request.Description,
			Link:        request.Link,
			Expired:     request.Expired,
			StartDate:   request.StartDate,
			Term:        request.Term,
		},
		HackathonStatusTag: request.StatusTags,
	}

	hackathon, err := server.store.CreateHackathonTx(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := HackathonResponses{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTags:  hackathon.HackathonStatusTag,
	}
	ctx.JSON(http.StatusOK, response)
}

// ハッカソン取得
// ハッカソンを取得する際のパラメータ
type GetHackathonParams struct {
	HackathonID int32 `uri:"hackathon_id"`
}

// ハッカソンを取得する
func (server *Server) GetHackathon(ctx *gin.Context) {
	var request GetHackathonParams
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hackathon, err := server.store.GetHackathon(ctx, request.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	statusTags, err := server.store.GetStatusTags(ctx, request.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var hackathonStatusTags []db.StatusTags

	for _, tags := range statusTags {
		statusTags, err := server.store.GetListStatusTags(ctx, tags.StatusID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		hackathonStatusTags = append(hackathonStatusTags, statusTags)
	}

	response := HackathonResponses{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTags:  hackathonStatusTags,
	}
	ctx.JSON(http.StatusOK, response)
}

// ハッカソン一覧取得
// ハッカソン一覧を取得する際のパラメータ
type ListHackathonsParams struct {
	PageSize int32 `form:"page_size"`
	PageId   int32 `form:"page_id"`
}

func (server *Server) ListHackathons(ctx *gin.Context) {
	var request ListHackathonsParams
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hackathons, err := server.store.ListHackathons(ctx, db.ListHackathonsParams{
		Expired: time.Now(),
		Limit:   request.PageSize,
		Offset:  (request.PageId - 1) * request.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	statusTags, err := server.store.ListHackathons(ctx, request.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var hackathonStatusTags []db.StatusTags

	for _, tags := range statusTags {
		statusTags, err := server.store.GetListStatusTags(ctx, tags.StatusID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		hackathonStatusTags = append(hackathonStatusTags, statusTags)
	}

	var response []HackathonResponses
	for _, hackathon := range hackathons {
		response = append(response, HackathonResponses{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon,
			Description: hackathon.Description,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			StatusTags:  hackathonStatusTags,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

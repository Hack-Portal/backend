package api

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

// ハッカソンを作る時のリクエストパラメータ
type CreateHackathonParams struct {
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
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
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	StatusTags []db.StatusTags `json:"status_tags"`
}

// ハッカソン作成
func (server *Server) CreateHackathon(ctx *gin.Context) {
	var (
		request  CreateHackathonParams
		imageURL string
	)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, _, err := ctx.Request.FormFile(ImageKey)
	if err != nil {
		switch err.Error() {
		case MultiPartNextPartEoF:
			ctx.JSON(400, errorResponse(err))
			return
		case HttpNoSuchFile:
			ctx.JSON(400, errorResponse(err))
			return
		default:
			ctx.JSON(400, errorResponse(err))
			return
		case RequestContentTypeIsnt:
			break
		}
	} else {
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			ctx.JSON(500, errorResponse(err))
			return
		}
		imageURL, err = server.store.UploadImage(ctx, buf.Bytes())
		if _, err := io.Copy(buf, file); err != nil {
			ctx.JSON(500, errorResponse(err))
			return
		}
	}

	args := db.CreateHackathonTxParams{
		Name:               request.Name,
		Icon:               imageURL,
		Description:        request.Description,
		Link:               request.Link,
		Expired:            request.Expired,
		StartDate:          request.StartDate,
		Term:               request.Term,
		HackathonStatusTag: request.StatusTags,
	}

	hackathon, err := server.store.CreateHackathonTx(ctx, &server.config, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := HackathonResponses{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
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

	hackathon, err := server.store.GetHackathonByID(ctx, request.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	statusTags, err := server.store.GetStatusTagsByHackathonID(ctx, request.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := HackathonResponses{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTags:  statusTags,
	}
	ctx.JSON(http.StatusOK, response)
}

// ハッカソン一覧取得
// ハッカソン一覧を取得する際のパラメータ
type ListHackathonsParams struct {
	PageSize int32 `form:"page_size"`
	PageId   int32 `form:"page_id"`
	Expired  bool  `form:"expired"`
}

type ListHackathonsResponses struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	StatusTags []db.StatusTags `json:"status_tags"`
}

func (server *Server) ListHackathons(ctx *gin.Context) {
	var request ListHackathonsParams
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var exp time.Time
	if request.Expired {
		exp = time.Now()
	} else {
		exp = time.Now().Add(-time.Hour * 720)
	}
	hackathons, err := server.store.ListHackathons(ctx, db.ListHackathonsParams{
		Expired: exp,
		Limit:   request.PageSize,
		Offset:  (request.PageId - 1) * request.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []ListHackathonsResponses

	for _, hackathon := range hackathons {
		statusTags, err := server.store.GetStatusTagsByHackathonID(ctx, hackathon.HackathonID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		response = append(response, ListHackathonsResponses{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			StatusTags:  statusTags,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

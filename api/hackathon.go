package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

type CreateHackathonParams struct {
	Name        string    `json:"Name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
}

type HackathonResponses struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
}

// ハッカソン作成
func (server *Server) CreateHackathon(ctx *gin.Context) {
	var request CreateHackathonParams
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateHackathonParams{
		Name: request.Name,
		Icon: sql.NullString{
			String: request.Icon,
			Valid:  true,
		},
		Description: request.Description,
		Link:        request.Link,
		Expired:     request.Expired,
		StartDate:   request.StartDate,
		Term:        request.Term,
	}

	hackathon, err := server.store.CreateHackathon(ctx, args)
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
	}
	ctx.JSON(http.StatusOK, response)
}

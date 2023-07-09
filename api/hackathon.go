package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

type CreateHackathonParams struct {
	Name        string    `json:"Name"`
	Icon        []byte    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
}

type HackathonResponses struct {
	Name        string    `json:"name"`
	Icon        []byte    `json:"icon"`
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := db.CreateHackathonParams{
		Name:        request.Name,
		Icon:        request.Icon,
		Description: request.Description,
		Link:        request.Link,
		Expired:     request.Expired,
		StartDate:   request.StartDate,
		Term:        request.Term,
	}

	hackathon, err := server.store.CreateHackathon(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := HackathonResponses{
		Name:        hackathon.Name,
		Icon:        hackathon.Icon,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
	}
	ctx.JSON(http.StatusOK, response)
}

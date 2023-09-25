package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type ChatController struct {
	Chatusecase *usecase.WsServer
	Env         *bootstrap.Env
}

func (cc *ChatController) ChatConnect(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI request.ChatRoom
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	cc.Chatusecase.ServeWs(ctx.Writer, ctx.Request, reqURI.RoomID, reqURI.AccountID)
}

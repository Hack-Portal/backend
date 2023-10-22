package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/ws"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type ChatController struct {
	Hub *ws.Hub
	Env *bootstrap.Env
	Db  transaction.Store
}

func (cc *ChatController) ChatRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI request.ChatRoom
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if _, err := cc.Db.GetRoomsByID(ctx, reqURI.RoomID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	members, err := cc.Db.GetRoomsAccountsByID(ctx, reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !checkMembers(members, reqURI.AccountID) {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ws.ServeWs(cc.Hub, cc.Db, ctx.Writer, ctx.Request, reqURI.AccountID, reqURI.RoomID)
}

func checkMembers(members []repository.GetRoomsAccountsByIDRow, accountID string) bool {
	for _, member := range members {
		if member.AccountID.String == accountID {
			return true
		}
	}
	return false
}

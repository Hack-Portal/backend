package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/ws"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

func NewRoomRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, hub *ws.Hub, group gin.IRoutes, publicGroup gin.IRoutes) {
	roomController := controller.RoomController{
		RoomUsecase: usecase.NewRoomUsercase(store, timeout),
		Env:         env,
	}

	chatController := controller.ChatController{
		Hub: hub,
		Env: env,
		Db:  store,
	}

	// ルーム
	group.GET("/rooms", roomController.ListRooms)
	group.POST("/rooms", roomController.CreateRoom)

	group.POST("/rooms/:room_id", roomController.AddAccountInRoom)
	group.GET("/rooms/:room_id", roomController.GetRoom)
	group.PUT("/rooms/:room_id", roomController.UpdateRoom)
	group.DELETE("/rooms/:room_id", roomController.DeleteRoom)

	group.POST("/rooms/:room_id/members", roomController.CloseRoom)
	group.DELETE("/rooms/:room_id/members", roomController.RemoveAccountInRoom)

	group.POST("/rooms/:room_id/addchat", roomController.AddChat)
	publicGroup.GET("/chats/:room_id/:account_id/", chatController.ChatRoom)

	group.POST("/rooms/:room_id/roles", roomController.AddRoomAccountRole)
	group.PUT("/rooms/:room_id/roles", roomController.UpdateRoomAccountRole)
}

package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

func NewRoomRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group gin.IRoutes) {
	roomController := controller.RoomController{
		RoomUsecase: usecase.NewRoomUsercase(store, timeout),
		Env:         env,
	}

	// ルーム
	group.GET("/rooms", roomController.ListRooms)
	group.POST("/rooms", roomController.CreateRoom)

	group.GET("/rooms/:room_id", roomController.GetRoom)
	group.POST("/rooms/:room_id", roomController.AddAccountInRoom)
	group.PUT("/rooms/:room_id", roomController.UpdateRoom)
	group.DELETE("/rooms/:room_id", roomController.DeleteRoom)

	group.POST("/rooms/:room_id/members", roomController.AddAccountInRoom)
	group.DELETE("/rooms/:room_id/members/user_id", roomController.RemoveAccountInRoom)
	group.POST("/rooms/:room_id/addchat", roomController.AddChat)
}

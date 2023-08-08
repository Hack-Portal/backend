package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/transaction"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

func NewRoomRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
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
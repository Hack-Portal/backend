package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

func NewRoomRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {

	roomRepository := repository.NewRoomRepository(store, domain.CollectionRoom)
	roomController := controller.RoomController{
		RoomUsecase: usecase.NewRoomUsercase(roomRepository, timeout),
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

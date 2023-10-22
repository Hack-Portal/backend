package inputport

import (
	"net/http"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain/entity"
)

type WsServer interface {
	findBrokerbyRoomID(ID string) *entity.Broker
	createBroker(room *entity.ChatRoom) *entity.Broker
	ServeWs(w http.ResponseWriter, req *http.Request, roomId string, userId string)
}

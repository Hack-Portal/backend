package usecase

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/entity"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		//origin := r.Header.Get("Origin")
		return true
	}}

type WsServer struct {
	Manager *BrokerManager
}

func (server *WsServer) findBrokerbyRoomID(ID string) *entity.Broker {
	for broker := range server.Manager.Brokers {
		if broker.Room.ID == ID {
			return broker
		}
	}
	return nil
}

func (server *WsServer) createBroker(room *entity.ChatRoom) *entity.Broker {
	broker := entity.NewBroker(room)
	go server.Manager.RunBroker(broker)
	server.Manager.Brokers[broker] = true

	return broker
}

func (server *WsServer) ServeWs(w http.ResponseWriter, req *http.Request, roomId string, userId string) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	if err != nil {
		return
	}
	room := &entity.ChatRoom{
		ID: roomId,
	}

	broker := server.findBrokerbyRoomID(room.ID)
	if broker == nil {
		broker = server.createBroker(room)
	}
	client := entity.NewClient(socket, userId, broker)
	clientManager := clientManager{
		client: client,
	}
	broker.Clients[client] = true
	broker.Join <- client
	defer func() { broker.Leave <- client }()
	go clientManager.clientWrite()
	clientManager.clientRead()
}

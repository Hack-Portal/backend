package entity

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

// struct representing a user in a ChatRoom
type Client struct {
	User         string
	LastActivity time.Time `json:"last_activity"`
	// The websocket Connection.
	Conn *websocket.Conn `json:"-"`
	// Buffered channel of outbound messages.
	Send chan *ChatEvent `json:"-"`
	// Broker for connection
	Broker *Broker
}

func NewClient(conn *websocket.Conn, userid string, broker *Broker) *Client {
	client := &Client{
		User:   userid,
		Conn:   conn,
		Send:   make(chan *ChatEvent, 100),
		Broker: broker,
	}
	return client
}

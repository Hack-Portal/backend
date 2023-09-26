package ws

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	hub *Hub

	accountID string
	roomID    string

	db transaction.Store

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Client) readPump() {
	ctx := context.Background()
	defer func() {
		c.hub.unregister <- c
		c.conn.Close(websocket.StatusNormalClosure, "")
	}()
	c.conn.SetReadLimit(maxMessageSize)
	for {
		_, message, err := c.conn.Read(ctx)
		if err != nil {
			// if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			// 	log.Printf("error: %v", err)
			// }
			break
		}

		if _, err = c.db.CreateChat(ctx, repository.CreateChatParams{
			ChatID:    uuid.New().String(),
			RoomID:    c.roomID,
			Message:   string(message),
			AccountID: c.accountID,
		}); err != nil {
			return
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	ctx := context.Background()
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close(websocket.StatusNormalClosure, "")
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The hub closed the channel.
				wsjson.Write(ctx, c.conn, []byte{})
				return
			}

			w, err := c.conn.Writer(ctx, websocket.MessageText)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := wsjson.Write(ctx, c.conn, []byte{}); err != nil {
				return
			}
		}
	}
}

func ServeWs(hub *Hub, db transaction.Store, w http.ResponseWriter, r *http.Request, accountID, roomID string) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true, OriginPatterns: []string{"*"}})
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, db: db, conn: conn, send: make(chan []byte, 256), accountID: accountID, roomID: roomID}
	client.hub.register <- client
	go client.writePump()
	client.readPump()
}

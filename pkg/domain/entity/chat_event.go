package entity

import (
	"time"
)

const (
	// Subscribe is used to broadcast a message indicating user has joined ChatRoom
	Subscribe = "join"
	// Broadcast is used to broadcast messages to all subscribed users
	Broadcast = "send"
	// Unsubscribe is used to broadcast a message indicating user has left ChatRoom
	Unsubscribe = "leave"
)

// struct representing a message event in an  ChatRoom
type ChatEvent struct {
	ID        string    `json:"_id"`
	EventType string    `json:"type"`
	UserID    string    `json:"user_id"`
	RoomID    string    `json:"room_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"time"`
}

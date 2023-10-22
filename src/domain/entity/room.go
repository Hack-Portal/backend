package entity

import (
	"time"
)

// struct representing a chat room
type ChatRoom struct {
	ID        string    `json:"_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Clients   []string  `json:"users"`
}

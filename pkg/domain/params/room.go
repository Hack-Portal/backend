package params

import "time"

type CreateRoom struct {
	RoomID      string `json:"room_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HackathonID int32  `json:"hackathon_id"`
	MemberLimit int32  `json:"member_limit"`
	OwnerID     string `json:"owner_id"`
	IncludeRate bool   `json:"include_rate"`
}

type UpdateRoom struct {
	RoomID      string `json:"room_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HackathonID int32  `json:"hackathon_id"`
	MemberLimit int32  `json:"member_limit"`
	IsClosing   bool   `json:"is_closing"`
	OwnerEmail  string `json:"owner_email"`
}

type DeleteRoom struct {
	OwnerEmail string `json:"owner_email"`
	RoomID     string `json:"room_id"`
}

type AddAccountInRoom struct {
	AccountID string `json:"account_id"`
	RoomID    string `json:"room_id"`
}

type AddChat struct {
	RoomID    string `json:"room_id"`
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
}

type WriteFireStore struct {
	RoomID  string `json:"room_id"`
	Index   int    `json:"index"`
	UID     string `json:"uid"`
	Message string `json:"message"`
}

type ChatRoomsWrite struct {
	UID       string    `json:"uid"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type DeleteRoomAccount struct {
	RoomID    string `json:"room_id"`
	Email     string `json:"email"`
	AccountID string `json:"account_id"`
}

type CloseRoom struct {
	RoomID    string
	AccountID []string
}

package params

import "time"

type CreateRoomParams struct {
	RoomID      string `json:"room_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HackathonID int32  `json:"hackathon_id"`
	MemberLimit int32  `json:"member_limit"`
	OwnerID     string `json:"owner_id"`
	IncludeRate bool   `json:"include_rate"`
}

type UpdateRoomParams struct {
	RoomID      string `json:"room_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HackathonID int32  `json:"hackathon_id"`
	MemberLimit int32  `json:"member_limit"`
	IsClosing   bool   `json:"is_closing"`
	OwnerEmail  string `json:"owner_email"`
}

type DeleteRoomParams struct {
	OwnerEmail string `json:"owner_email"`
	RoomID     string `json:"room_id"`
}

type AddAccountInRoomParams struct {
	AccountID string `json:"account_id"`
	RoomID    string `json:"room_id"`
}

type AddChatParams struct {
	RoomID    string `json:"room_id"`
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
}

type WriteFireStoreParams struct {
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

type CloseRoomParams struct {
	RoomID    string
	AccountID []string
}

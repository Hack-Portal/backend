package request

type RoomsWildCard struct {
	RoomID string `uri:"room_id"`
}

type CreateRoom struct {
	HackathonID int32  `json:"hackathon_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	MemberLimit int32  `json:"member_limit" binding:"required"`
	AccountID   string `json:"account_id" binding:"required"`
}

type AddChat struct {
	AccountID string `json:"account_id" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

type UpdateRoom struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MemberLimit int32  `json:"member_limit"`
	HackathonID int32  `json:"hackathonID"`
	IsClosing   bool   `json:"is_closing"`
}

type AddAccountInRoom struct {
	AccountID string `json:"account_id"`
}

type RemoveAccountInRoom struct {
	AccountID string `form:"account_id"`
}

type RoomAccountRole struct {
	AccountID string  `json:"account_id"`
	RoleID    []int32 `json:"role_id"`
}

type CloseRoom struct {
	AccountID []string `json:"account_id"`
}

type ChatRoom struct {
	AccountID string `uri:"account_id"`
	RoomID    string `uri:"room_id"`
}

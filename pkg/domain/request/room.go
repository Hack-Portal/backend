package request

type RoomsRequestWildCard struct {
	RoomID string `uri:"room_id"`
}

type CreateRoomRequestBody struct {
	HackathonID int32  `json:"hackathon_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	MemberLimit int32  `json:"member_limit" binding:"required"`
	AccountID   string `json:"account_id" binding:"required"`
}

type AddChatRequestBody struct {
	AccountID string `json:"account_id" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

type UpdateRoomRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MemberLimit int32  `json:"member_limit"`
	HackathonID int32  `json:"hackathonID"`
	IsClosing   bool   `json:"is_closing"`
}

type AddAccountInRoomRequestBody struct {
	AccountID string `json:"account_id"`
}

type RemoveAccountInRoomRequest struct {
	AccountID string `form:"account_id"`
}

type CloseRoomRequest struct {
	AccountID []string `json:"account_id"`
}

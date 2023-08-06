package domain

import (
	"time"

	"github.com/gofrs/uuid"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

type RoomsRequestWildCard struct {
	RoomID string `uri:"room_id"`
}

type CreateRoomRequestBody struct {
	HackathonID int32  `json:"hackathon_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	MemberLimit int32  `json:"member_limit" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
}

type ListRoomsRequest struct {
	PageSize int32 `form:"page_size"`
	PageID   int32 `from:"page_id"`
}

type RoomTechTags struct {
	TechTag db.TechTags `json:"tech_tag"`
	Count   int32       `json:"count"`
}
type RoomFramework struct {
	Framework db.Frameworks `json:"framework"`
	Count     int32         `json:"count"`
}

type NowRoomAccounts struct {
	UserID  string `json:"user_id"`
	Icon    string `json:"icon"`
	IsOwner bool   `json:"is_owner"`
}

type ListRoomRoomInfo struct {
	RoomID      uuid.UUID `json:"room_id"`
	Title       string    `json:"title"`
	MemberLimit int32     `json:"member_limit"`
	CreatedAt   time.Time `json:"created_at"`
}
type ListRoomHackathonInfo struct {
	HackathonID   int32  `json:"hackathon_id"`
	HackathonName string `json:"hackathon_name"`
	Icon          string `json:"icon"`
}
type ListRoomResponse struct {
	Rooms             ListRoomRoomInfo      `json:"rooms"`
	Hackathon         ListRoomHackathonInfo `json:"hackathon"`
	NowMember         []NowRoomAccounts     `json:"now_member"`
	MembersTechTags   []RoomTechTags        `json:"members_tech_tags"`
	MembersFrameworks []RoomFramework       `json:"members_frameworks"`
}

type hackathonInfo struct {
	HackathonID int32           `json:"hackathon_id"`
	Name        string          `json:"name"`
	Icon        string          `json:"icon"`
	Description string          `json:"description"`
	Link        string          `json:"link"`
	Expired     time.Time       `json:"expired"`
	StartDate   time.Time       `json:"start_date"`
	Term        int32           `json:"term"`
	Tags        []db.StatusTags `json:"tags"`
}

type GetRoomResponse struct {
	RoomID            uuid.UUID            `json:"room_id"`
	Title             string               `json:"title"`
	Description       string               `json:"description"`
	MemberLimit       int32                `json:"member_limit"`
	IsStatus          bool                 `json:"is_status"`
	CreateAt          time.Time            `json:"create_at"`
	Hackathon         hackathonInfo        `json:"hackathon"`
	NowMember         []db.NowRoomAccounts `json:"now_member"`
	MembersTechTags   []db.RoomTechTags    `json:"members_tech_tags"`
	MembersFrameworks []db.RoomFramework   `json:"members_frameworks"`
}

type AddChatRequestBody struct {
	Message string `json:"message" binding:"required"`
}

type UpdateRoomRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	MemberLimit int32  `json:"member_limit" binding:"required"`
}

type RoomUsecase interface {
	ListRooms
	CreateRoom
	GetRoom
	AddAccountInRoom
	UpdateRoom
	DeleteRoom
	AddAccountIn
}

package domain

import (
	"time"

	"github.com/google/uuid"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
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
	TechTag repository.TechTag `json:"tech_tag"`
	Count   int32              `json:"count"`
}
type RoomFramework struct {
	Framework repository.Framework `json:"framework"`
	Count     int32                `json:"count"`
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

type HackathonInfo struct {
	HackathonID int32                  `json:"hackathon_id"`
	Name        string                 `json:"name"`
	Icon        string                 `json:"icon"`
	Link        string                 `json:"link"`
	StartDate   time.Time              `json:"start_date"`
	Term        int32                  `json:"term"`
	Tags        []repository.StatusTag `json:"tags"`
}

type GetRoomResponse struct {
	RoomID            uuid.UUID         `json:"room_id"`
	Title             string            `json:"title"`
	Description       string            `json:"description"`
	MemberLimit       int32             `json:"member_limit"`
	IsDelete          bool              `json:"is_status"`
	CreateAt          time.Time         `json:"create_at"`
	Hackathon         HackathonInfo     `json:"hackathon"`
	NowMember         []NowRoomAccounts `json:"now_member"`
	MembersTechTags   []RoomTechTags    `json:"members_tech_tags"`
	MembersFrameworks []RoomFramework   `json:"members_frameworks"`
}

type AddChatRequestBody struct {
	UserID  string `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type UpdateRoomRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MemberLimit int32  `json:"member_limit"`
	HackathonID int32  `json:"hackathonID"`
}

type CreateRoomParam struct {
	RoomID      uuid.UUID `json:"room_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	HackathonID int32     `json:"hackathon_id"`
	MemberLimit int32     `json:"member_limit"`
	OwnerID     string    `json:"owner_id"`
}

type UpdateRoomParam struct {
	RoomID      uuid.UUID `json:"room_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	HackathonID int32     `json:"hackathon_id"`
	MemberLimit int32     `json:"member_limit"`
	OwnerEmail  string    `json:"owner_email"`
}

type DeleteRoomParam struct {
	OwnerEmail string    `json:"owner_email"`
	RoomID     uuid.UUID `json:"room_id"`
}

type AddAccountInRoomParam struct {
	UserID string    `json:"user_id"`
	RoomID uuid.UUID `json:"room_id"`
}

type AddAccountInRoomRequestBody struct {
	UserID string `json:"user_id"`
}

type AddChatParams struct {
	RoomID  uuid.UUID `json:"room_id"`
	UserID  string    `json:"user_id"`
	Message string    `json:"message"`
}

type WriteFireStoreParam struct {
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

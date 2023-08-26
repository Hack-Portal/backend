package domain

import (
	"time"

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
	AccountID   string `json:"account_id" binding:"required"`
}
type RoomTechTags struct {
	TechTag repository.TechTag `json:"tech_tag"`
	Count   int32              `json:"count"`
}
type RoomFramework struct {
	Framework repository.Framework `json:"framework"`
	Count     int32                `json:"count"`
}
type RoomRole struct {
	Role repository.Role `json:"role"`
}
type NowRoomAccounts struct {
	AccountID  string                 `json:"account_id"`
	Username   string                 `json:"username"`
	Icon       string                 `json:"icon"`
	IsOwner    bool                   `json:"is_owner"`
	Roles      []repository.Role      `json:"roles"`
	TechTags   []repository.TechTag   `json:"tech_tags"`
	Frameworks []repository.Framework `json:"frameworks"`
}

type ListRoomRoomInfo struct {
	RoomID      string    `json:"room_id"`
	Title       string    `json:"title"`
	MemberLimit int32     `json:"member_limit"`
	CreatedAt   time.Time `json:"created_at"`
}
type ListRoomHackathonInfo struct {
	HackathonID   int32     `json:"hackathon_id"`
	HackathonName string    `json:"hackathon_name"`
	Icon          string    `json:"icon"`
	Expired       time.Time `json:"expired"`
}
type ListRoomResponse struct {
	Rooms             ListRoomRoomInfo      `json:"rooms"`
	Hackathon         ListRoomHackathonInfo `json:"hackathon"`
	NowMember         []NowRoomAccounts     `json:"now_member"`
	MembersTechTags   []RoomTechTags        `json:"members_tech_tags"`
	MembersFrameworks []RoomFramework       `json:"members_frameworks"`
}

type RoomHackathonInfo struct {
	HackathonID int32                  `json:"hackathon_id"`
	Name        string                 `json:"name"`
	Icon        string                 `json:"icon"`
	Link        string                 `json:"link"`
	StartDate   time.Time              `json:"start_date"`
	Expired     time.Time              `json:"expired"`
	Term        int32                  `json:"term"`
	StatusTag   []repository.StatusTag `json:"status_tag"`
}

type GetRoomResponse struct {
	RoomID      string            `json:"room_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	MemberLimit int32             `json:"member_limit"`
	IsDelete    bool              `json:"is_status"`
	Hackathon   RoomHackathonInfo `json:"hackathon"`
	NowMember   []NowRoomAccounts `json:"now_member"`
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
}

type CreateRoomParam struct {
	RoomID      string `json:"room_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HackathonID int32  `json:"hackathon_id"`
	MemberLimit int32  `json:"member_limit"`
	OwnerID     string `json:"owner_id"`
	IncludeRate bool   `json:"include_rate"`
}

type UpdateRoomParam struct {
	RoomID      string `json:"room_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HackathonID int32  `json:"hackathon_id"`
	MemberLimit int32  `json:"member_limit"`
	OwnerEmail  string `json:"owner_email"`
}

type DeleteRoomParam struct {
	OwnerEmail string `json:"owner_email"`
	RoomID     string `json:"room_id"`
}

type AddAccountInRoomParam struct {
	AccountID string `json:"account_id"`
	RoomID    string `json:"room_id"`
}

type AddAccountInRoomRequestBody struct {
	AccountID string `json:"account_id"`
}

type AddChatParams struct {
	RoomID    string `json:"room_id"`
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
}

type RoomAccountRoleByIDParam struct {
	RoomID    string `json:"room_id"`
	AccountID string `json:"account_id"`
	RoleID    int32  `json:"role_id"`
}

type RoomAccountRoleByIDRequestBody struct {
	AccountID string `json:"account_id"`
	RoleID    int32  `json:"role_id"`
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

type DeleteRoomAccount struct {
	RoomID    string `json:"room_id"`
	Email     string `json:"email"`
	AccountID string `json:"account_id"`
}

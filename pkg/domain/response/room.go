package response

import (
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type ListRoomRoomInfo struct {
	RoomID      string    `json:"room_id"`
	Title       string    `json:"title"`
	MemberLimit int32     `json:"member_limit"`
	IsClosing   bool      `json:"is_closing"`
	CreatedAt   time.Time `json:"created_at"`
}
type ListRoomHackathonInfo struct {
	HackathonID   int32     `json:"hackathon_id"`
	HackathonName string    `json:"hackathon_name"`
	Icon          string    `json:"icon"`
	Expired       time.Time `json:"expired"`
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

type RoomTechTags struct {
	TechTag repository.TechTag `json:"tech_tag"`
	Count   int32              `json:"count"`
}
type RoomFramework struct {
	Framework repository.Framework `json:"framework"`
	Count     int32                `json:"count"`
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

type ListRoomResponse struct {
	Rooms             ListRoomRoomInfo      `json:"rooms"`
	Hackathon         ListRoomHackathonInfo `json:"hackathon"`
	NowMember         []NowRoomAccounts     `json:"now_member"`
	MembersTechTags   []RoomTechTags        `json:"members_tech_tags"`
	MembersFrameworks []RoomFramework       `json:"members_frameworks"`
}

type GetRoomResponse struct {
	RoomID      string            `json:"room_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	MemberLimit int32             `json:"member_limit"`
	IsClosing   bool              `json:"is_closing"`
	IsDelete    bool              `json:"is_status"`
	Hackathon   RoomHackathonInfo `json:"hackathon"`
	NowMember   []NowRoomAccounts `json:"now_member"`
}

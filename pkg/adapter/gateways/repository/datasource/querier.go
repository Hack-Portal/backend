// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package repository

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccountFrameworks(ctx context.Context, arg CreateAccountFrameworksParams) (AccountFramework, error)
	CreateAccountPastWorks(ctx context.Context, arg CreateAccountPastWorksParams) (AccountPastWork, error)
	CreateAccountTags(ctx context.Context, arg CreateAccountTagsParams) (AccountTag, error)
	CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error)
	CreateBookmarks(ctx context.Context, arg CreateBookmarksParams) (Bookmark, error)
	CreateFollows(ctx context.Context, arg CreateFollowsParams) (Follow, error)
	CreateHackathonStatusTags(ctx context.Context, arg CreateHackathonStatusTagsParams) (HackathonStatusTag, error)
	CreateHackathons(ctx context.Context, arg CreateHackathonsParams) (Hackathon, error)
	CreatePastWorkFrameworks(ctx context.Context, arg CreatePastWorkFrameworksParams) (PastWorkFramework, error)
	CreatePastWorkTags(ctx context.Context, arg CreatePastWorkTagsParams) (PastWorkTag, error)
	CreatePastWorks(ctx context.Context, arg CreatePastWorksParams) (PastWork, error)
	CreateRateEntries(ctx context.Context, arg CreateRateEntriesParams) (RateEntry, error)
	CreateRooms(ctx context.Context, arg CreateRoomsParams) (Room, error)
	CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccount, error)
	DeleteAccountFrameworkByUserID(ctx context.Context, userID string) error
	DeleteAccountPastWorksByOpus(ctx context.Context, opus int32) error
	DeleteAccountTagsByUserID(ctx context.Context, userID string) error
	DeleteAccounts(ctx context.Context, userID string) (Account, error)
	DeleteBookmarksByID(ctx context.Context, arg DeleteBookmarksByIDParams) (Bookmark, error)
	DeleteFollows(ctx context.Context, arg DeleteFollowsParams) error
	DeleteFrameworksByID(ctx context.Context, frameworkID int32) error
	DeleteHackathonByID(ctx context.Context, hackathonID int32) error
	DeleteHackathonStatusTagsByID(ctx context.Context, hackathonID int32) error
	DeletePastWorkFrameworksByOpus(ctx context.Context, opus int32) error
	DeletePastWorkTagsByOpus(ctx context.Context, opus int32) error
	DeleteRoomsAccountsByID(ctx context.Context, arg DeleteRoomsAccountsByIDParams) error
	DeleteRoomsByID(ctx context.Context, roomID uuid.UUID) (Room, error)
	DeleteStatusTagsByStatusID(ctx context.Context, statusID int32) error
	DeleteTechTagsByID(ctx context.Context, techTagID int32) error
	GetAccountsByEmail(ctx context.Context, email string) (GetAccountsByEmailRow, error)
	GetAccountsByID(ctx context.Context, userID string) (GetAccountsByIDRow, error)
	GetFrameworksByID(ctx context.Context, frameworkID int32) (Framework, error)
	GetHackathonByID(ctx context.Context, hackathonID int32) (Hackathon, error)
	GetLocatesByID(ctx context.Context, locateID int32) (Locate, error)
	GetPastWorksByOpus(ctx context.Context, opus int32) (PastWork, error)
	GetRoomsAccountsByID(ctx context.Context, roomID uuid.UUID) ([]GetRoomsAccountsByIDRow, error)
	GetRoomsByID(ctx context.Context, roomID uuid.UUID) (Room, error)
	GetStatusTagsByHackathonID(ctx context.Context, hackathonID int32) ([]StatusTag, error)
	GetStatusTagsByStatusID(ctx context.Context, statusID int32) (StatusTag, error)
	GetTechTagsByID(ctx context.Context, techTagID int32) (TechTag, error)
	ListAccountFrameworksByUserID(ctx context.Context, userID string) ([]ListAccountFrameworksByUserIDRow, error)
	ListAccountPastWorksByOpus(ctx context.Context, opus int32) ([]AccountPastWork, error)
	ListAccountTagsByUserID(ctx context.Context, userID string) ([]ListAccountTagsByUserIDRow, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]ListAccountsRow, error)
	ListBookmarksByID(ctx context.Context, userID string) ([]Bookmark, error)
	ListFollowsByToUserID(ctx context.Context, toUserID string) ([]Follow, error)
	ListFrameworks(ctx context.Context) ([]Framework, error)
	ListHackathonStatusTagsByID(ctx context.Context, hackathonID int32) ([]HackathonStatusTag, error)
	ListHackathons(ctx context.Context, arg ListHackathonsParams) ([]Hackathon, error)
	ListLocates(ctx context.Context) ([]Locate, error)
	ListPastWorkFrameworksByOpus(ctx context.Context, opus int32) ([]PastWorkFramework, error)
	ListPastWorkTagsByOpus(ctx context.Context, opus int32) ([]PastWorkTag, error)
	ListPastWorks(ctx context.Context, arg ListPastWorksParams) ([]ListPastWorksRow, error)
	ListRateEntries(ctx context.Context, arg ListRateEntriesParams) ([]RateEntry, error)
	ListRooms(ctx context.Context, arg ListRoomsParams) ([]Room, error)
	ListStatusTags(ctx context.Context) ([]StatusTag, error)
	ListTechTags(ctx context.Context) ([]TechTag, error)
	UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Account, error)
	UpdateFrameworksByID(ctx context.Context, arg UpdateFrameworksByIDParams) (Framework, error)
	UpdateRateByID(ctx context.Context, arg UpdateRateByIDParams) (Account, error)
	UpdateRoomsByID(ctx context.Context, arg UpdateRoomsByIDParams) (Room, error)
	UpdateStatusTagsByStatusID(ctx context.Context, status string) (StatusTag, error)
	UpdateTechTagsByID(ctx context.Context, language string) (TechTag, error)
}

var _ Querier = (*Queries)(nil)
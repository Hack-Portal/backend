// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreataAccountTags(ctx context.Context, arg CreataAccountTagsParams) (AccountTags, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateAccountFramework(ctx context.Context, arg CreateAccountFrameworkParams) (AccountFrameworks, error)
	CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmarks, error)
	CreateFollow(ctx context.Context, arg CreateFollowParams) (Follows, error)
	CreateHackathon(ctx context.Context, arg CreateHackathonParams) (Hackathons, error)
	CreateHackathonStatusTag(ctx context.Context, arg CreateHackathonStatusTagParams) (HackathonStatusTags, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Rooms, error)
	CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccounts, error)
	GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error)
	GetAccountByID(ctx context.Context, userID string) (GetAccountByIDRow, error)
	GetFrameworksByID(ctx context.Context, frameworkID int32) (Frameworks, error)
	GetHackathonByID(ctx context.Context, hackathonID int32) (Hackathons, error)
	GetHackathonStatusTagsByHackathonID(ctx context.Context, hackathonID int32) ([]HackathonStatusTags, error)
	GetLocateByID(ctx context.Context, locateID int32) (Locates, error)
	GetRoomsAccountsByRoomID(ctx context.Context, roomID uuid.UUID) ([]GetRoomsAccountsByRoomIDRow, error)
	GetRoomsByID(ctx context.Context, roomID uuid.UUID) (Rooms, error)
	GetStatusTagByStatusID(ctx context.Context, statusID int32) (StatusTags, error)
	GetStatusTagsByHackathonID(ctx context.Context, hackathonID int32) ([]StatusTags, error)
	GetTechTagByID(ctx context.Context, techTagID int32) (TechTags, error)
	ListAccountFrameworksByUserID(ctx context.Context, userID string) ([]ListAccountFrameworksByUserIDRow, error)
	ListAccountTagsByUserID(ctx context.Context, userID string) ([]ListAccountTagsByUserIDRow, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]ListAccountsRow, error)
	ListBookmarkByUserID(ctx context.Context, userID string) ([]Bookmarks, error)
	ListFollowByToUserID(ctx context.Context, toUserID string) ([]Follows, error)
	ListFrameworks(ctx context.Context, limit int32) ([]Frameworks, error)
	ListHackathons(ctx context.Context, arg ListHackathonsParams) ([]Hackathons, error)
	ListLocates(ctx context.Context) ([]Locates, error)
	ListRoom(ctx context.Context, limit int32) ([]Rooms, error)
	ListStatusTags(ctx context.Context) ([]StatusTags, error)
	ListTechTag(ctx context.Context) ([]TechTags, error)
	RemoveAccountInRoom(ctx context.Context, arg RemoveAccountInRoomParams) error
	RemoveBookmark(ctx context.Context, arg RemoveBookmarkParams) error
	RemoveFollow(ctx context.Context, arg RemoveFollowParams) error
}

var _ Querier = (*Queries)(nil)

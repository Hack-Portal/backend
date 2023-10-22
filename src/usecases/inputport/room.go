package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
)

type RoomUsecase interface {
	ListRooms(ctx context.Context, query request.ListRequest) ([]response.ListRoom, error)
	GetRoom(ctx context.Context, id string) (result response.Room, err error)
	CreateRoom(ctx context.Context, body params.CreateRoom) (result response.Room, err error)
	UpdateRoom(ctx context.Context, body params.UpdateRoom) (result response.Room, err error)
	DeleteRoom(ctx context.Context, query params.DeleteRoom) error
	AddAccountInRoom(ctx context.Context, query params.AddAccountInRoom) error
	AddChat(ctx context.Context, body params.AddChat) error
	DeleteRoomAccount(ctx context.Context, body params.DeleteRoomAccount) (err error)
	UpdateRoomAccountRole(ctx context.Context, body params.RoomAccountRole) (err error)
	AddRoomAccountRole(ctx context.Context, body params.RoomAccountRole) error
	CloseRoom(ctx context.Context, body params.CloseRoom) error
}

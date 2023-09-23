package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
)

type RoomUsecase interface {
	ListRooms(ctx context.Context, query request.ListRequest) ([]response.ListRoomResponse, error)
	GetRoom(ctx context.Context, id string) (result response.GetRoomResponse, err error)
	CreateRoom(ctx context.Context, body params.CreateRoom) (result response.GetRoomResponse, err error)
	UpdateRoom(ctx context.Context, body params.UpdateRoom) (result response.GetRoomResponse, err error)
	DeleteRoom(ctx context.Context, query params.DeleteRoom) error
	AddAccountInRoom(ctx context.Context, query params.AddAccountInRoom) error
	AddChat(ctx context.Context, body params.AddChat) error
	DeleteRoomAccount(ctx context.Context, body params.DeleteRoomAccount) (err error)
	CloseRoom(ctx context.Context, body params.CloseRoom) error
}

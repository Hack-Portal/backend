package inputport

import (
	"context"

	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type RoomUsecase interface {
	ListRooms(ctx context.Context, query domain.ListRoomsRequest) (result []domain.ListRoomResponse, err error)
	GetRoom(ctx context.Context, id uuid.UUID) (result domain.GetRoomResponse, err error)
	CreateRoom(ctx context.Context, body domain.CreateRoomParam) (result domain.GetRoomResponse, err error)
	UpdateRoom(ctx context.Context, body domain.UpdateRoomParam) (result domain.GetRoomResponse, err error)
	DeleteRoom(ctx context.Context, query domain.DeleteRoomParam) error
	AddAccountInRoom(ctx context.Context, query domain.AddAccountInRoomParam) error
	AddChat(ctx context.Context, body domain.AddChatParams) error
}

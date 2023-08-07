package inputport

import (
	"context"

	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type RoomUsecase interface {
	ListRooms(ctx context.Context, query domain.ListRoomsRequest) (result []domain.ListRoomResponse, err error)
	CreateRoom(ctx context.Context, id uuid.UUID) (result domain.GetRoomResponse, err error)
	GetRoom
	AddAccountInRoom
	UpdateRoom
	DeleteRoom
	AddAccountIn
}

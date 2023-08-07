package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/domain"
	"github.com/hackhack-Geek-vol6/backend/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/usecase/inputport"
)

type roomUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewRoomUsercase(store transaction.Store, timeout time.Duration) inputport.RoomUsecase {
	return &roomUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (ru *roomUsecase) ListRooms(ctx context.Context, query domain.ListRoomsRequest) (result []domain.ListRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	return ru.store.ListRoomTx(ctx, query)
}

func (ru *roomUsecase) GetRoom(ctx context.Context, id uuid.UUID) (result domain.GetRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	return ru.store.GetRoomTx(ctx, id)
}

func (ru *roomUsecase) CreateRoom(ctx context.Context, body domain.CreateRoomParam) (result domain.GetRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	body.RoomID = uuid.New()
	// チャットルームの初期化
	_, err = ru.store.InitChatRoom(ctx, body.RoomID.String())
	if err != nil {
		return domain.GetRoomResponse{}, err
	}

	return ru.store.CreateRoomTx(ctx, body)
}

func (ru *roomUsecase) UpdateRoom(ctx context.Context, body domain.UpdateRoomParam) (result domain.GetRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	return ru.store.UpdateRoomTx(ctx, body)
}

func (ru *roomUsecase) DeleteRoom(ctx context.Context, query domain.DeleteRoomParam) error {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	return ru.store.DeleteRoomTx(ctx, query)
}

func (ru *roomUsecase) AddAccountInRoom(ctx context.Context, query domain.AddAccountInRoomParam) error {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	return ru.store.AddAccountInRoom(ctx, query)
}

func (ru *roomUsecase) AddChat(ctx context.Context, body domain.AddChatParams) error {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	data, err := ru.store.ReadDocsByRoomID(ctx, body.RoomID.String())
	if err != nil {
		return err
	}
	_, err = ru.store.WriteFireStore(ctx, domain.WriteFireStoreParam{
		RoomID:  body.RoomID.String(),
		Index:   len(data) + 1,
		UID:     body.UserID,
		Message: body.Message,
	})

	return err
}

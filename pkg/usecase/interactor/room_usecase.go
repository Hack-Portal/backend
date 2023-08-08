package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
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

func stackTagAndFrameworks(ctx context.Context, q *repository.Queries, room repository.Room) ([]domain.RoomTechTags, []domain.RoomFramework, error) {
	var (
		roomTechTags   []domain.RoomTechTags
		roomFrameworks []domain.RoomFramework
	)
	accounts, err := q.GetRoomsAccountsByID(ctx, room.RoomID)
	if err != nil {
		return nil, nil, err
	}

	for _, account := range accounts {
		techTags, err := q.ListAccountTagsByUserID(ctx, account.UserID.String)
		if err != nil {
			return nil, nil, err
		}
		for _, techTag := range techTags {
			roomTechTags = margeTechTagArray(roomTechTags, repository.TechTag{
				TechTagID: techTag.TechTagID.Int32,
				Language:  techTag.Language.String,
			})
		}

		frameworks, err := q.ListAccountFrameworksByUserID(ctx, account.UserID.String)
		if err != nil {
			return nil, nil, err
		}
		for _, framework := range frameworks {
			roomFrameworks = margeFrameworkArray(roomFrameworks, repository.Framework{
				FrameworkID: framework.FrameworkID.Int32,
				TechTagID:   framework.TechTagID.Int32,
				Framework:   framework.Framework.String,
			})
		}
	}
	return roomTechTags, roomFrameworks, nil
}

func margeTechTagArray(roomTechTags []domain.RoomTechTags, techtag repository.TechTag) []domain.RoomTechTags {
	for _, roomTechTag := range roomTechTags {
		if roomTechTag.TechTag == techtag {
			roomTechTag.Count++
		}
	}
	roomTechTags = append(roomTechTags, domain.RoomTechTags{
		TechTag: techtag,
		Count:   1,
	})

	return roomTechTags
}

func margeFrameworkArray(roomFramework []domain.RoomFramework, framework repository.Framework) []domain.RoomFramework {
	for _, roomFramework := range roomFramework {
		if roomFramework.Framework == framework {
			roomFramework.Count++
		}
	}
	roomFramework = append(roomFramework, domain.RoomFramework{
		Framework: framework,
		Count:     1,
	})

	return roomFramework
}

func margeRoomAccount(ctx context.Context, q *repository.Queries, id uuid.UUID) (result []domain.NowRoomAccounts, err error) {
	nowMembers, err := q.GetRoomsAccountsByID(ctx, id)
	if err != nil {
		return
	}

	for _, nowMember := range nowMembers {
		result = append(result, domain.NowRoomAccounts{
			UserID:  nowMember.UserID.String,
			Icon:    nowMember.Icon.String,
			IsOwner: nowMember.IsOwner,
		})
	}
	return
}

func parseRoomResponse(response domain.GetRoomResponse, room repository.Room, hackathon domain.HackathonInfo) domain.GetRoomResponse {
	return domain.GetRoomResponse{
		RoomID:      room.RoomID,
		Title:       room.Title,
		Description: room.Description,
		MemberLimit: room.MemberLimit,
		IsDelete:    room.IsDelete,
		CreateAt:    room.CreateAt,
		Hackathon:   hackathon,
	}
}

func getRoomMember(ctx context.Context, q *repository.Queries, id uuid.UUID) (result []domain.NowRoomAccounts, err error) {
	accounts, err := q.GetRoomsAccountsByID(ctx, id)
	if err != nil {
		return
	}

	for _, account := range accounts {
		user, err := q.GetAccountsByID(ctx, account.UserID.String)
		if err != nil {
			return nil, err
		}
		result = append(result, domain.NowRoomAccounts{
			UserID:  user.UserID,
			Icon:    user.Icon.String,
			IsOwner: account.IsOwner,
		})
	}
	return
}

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

func (ru *roomUsecase) ListRooms(ctx context.Context, query domain.ListRequest) ([]domain.ListRoomResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	var result []domain.ListRoomResponse

	rooms, err := ru.store.ListRooms(ctx, repository.ListRoomsParams{Limit: query.PageSize, Offset: (query.PageID - 1) * query.PageSize})
	if err != nil {
		return nil, err
	}

	for _, room := range rooms {
		hackathon, err := ru.store.GetHackathonByID(ctx, room.HackathonID)
		if err != nil {
			return nil, err
		}
		techtags, frameworks, err := stackTagAndFrameworks(ctx, ru.store, room)
		if err != nil {
			return nil, err
		}

		members, err := getRoomMember(ctx, ru.store, room.RoomID)
		if err != nil {
			return nil, err
		}
		result = append(result, domain.ListRoomResponse{
			Rooms: domain.ListRoomRoomInfo{
				RoomID:      room.RoomID,
				Title:       room.Title,
				MemberLimit: room.MemberLimit,
				IsClosing:   room.IsClosing.Bool,
				CreatedAt:   room.CreateAt,
			},
			Hackathon: domain.ListRoomHackathonInfo{
				HackathonID:   hackathon.HackathonID,
				HackathonName: hackathon.Name,
				Icon:          hackathon.Icon.String,
				Expired:       hackathon.Expired,
			},
			NowMember:         members,
			MembersTechTags:   techtags,
			MembersFrameworks: frameworks,
		})
	}
	return result, nil
}

func (ru *roomUsecase) GetRoom(ctx context.Context, id string) (result domain.GetRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	room, err := ru.store.GetRoomsByID(ctx, id)
	if err != nil {
		return
	}

	hackathon, err := ru.store.GetHackathonByID(ctx, room.HackathonID)
	if err != nil {
		return
	}

	statusTag, err := getHackathonTag(ctx, ru.store, hackathon.HackathonID)
	if err != nil {
		return
	}

	result.NowMember, err = getRoomMember(ctx, ru.store, id)
	if err != nil {
		return
	}

	result = parseRoomResponse(result, room, domain.RoomHackathonInfo{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Link:        hackathon.Link,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTag:   statusTag,
		Expired:     hackathon.Expired,
	})
	return
}

func (ru *roomUsecase) CreateRoom(ctx context.Context, body domain.CreateRoomParam) (result domain.GetRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	body.RoomID = uuid.New().String()

	room, err := ru.store.CreateRoomTx(ctx, body)
	if err != nil {
		return
	}

	hackathon, err := ru.store.GetHackathonByID(ctx, room.HackathonID)
	if err != nil {
		return
	}

	statusTag, err := getHackathonTag(ctx, ru.store, hackathon.HackathonID)
	if err != nil {
		return
	}

	result.NowMember, err = getRoomMember(ctx, ru.store, room.RoomID)
	if err != nil {
		return
	}

	result = parseRoomResponse(result, room, domain.RoomHackathonInfo{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Link:        hackathon.Link,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTag:   statusTag,
		Expired:     hackathon.Expired,
	})

	return
}

func (ru *roomUsecase) UpdateRoom(ctx context.Context, body domain.UpdateRoomParam) (result domain.GetRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	room, err := ru.store.UpdateRoomTx(ctx, body)
	if err != nil {
		return
	}
	hackathon, err := ru.store.GetHackathonByID(ctx, room.HackathonID)
	if err != nil {
		return
	}

	statusTag, err := getHackathonTag(ctx, ru.store, hackathon.HackathonID)
	if err != nil {
		return
	}

	result.NowMember, err = getRoomMember(ctx, ru.store, room.RoomID)
	if err != nil {
		return
	}

	result = parseRoomResponse(result, room, domain.RoomHackathonInfo{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Link:        hackathon.Link,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTag:   statusTag,
	})

	return
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

	data, err := ru.store.ReadDocsByRoomID(ctx, body.RoomID)
	if err != nil {
		return err
	}
	_, err = ru.store.CreateSubCollection(ctx, domain.WriteFireStoreParam{
		RoomID:  body.RoomID,
		Index:   data + 1,
		UID:     body.AccountID,
		Message: body.Message,
	})

	return err
}

func (ru *roomUsecase) DeleteRoomAccount(ctx context.Context, body domain.DeleteRoomAccount) (err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	account, err := ru.store.GetAccountsByEmail(ctx, body.Email)
	if err != nil {
		return
	}

	err = ru.store.DeleteRoomsAccountsByID(ctx, repository.DeleteRoomsAccountsByIDParams{
		RoomID:    body.RoomID,
		AccountID: account.AccountID,
	})
	return
}

func stackTagAndFrameworks(ctx context.Context, store transaction.Store, room repository.Room) ([]domain.RoomTechTags, []domain.RoomFramework, error) {
	var (
		roomTechTags   []domain.RoomTechTags
		roomFrameworks []domain.RoomFramework
	)
	accounts, err := store.GetRoomsAccountsByID(ctx, room.RoomID)
	if err != nil {
		return nil, nil, err
	}

	for _, account := range accounts {
		techTags, err := store.ListAccountTagsByUserID(ctx, account.AccountID.String)
		if err != nil {
			return nil, nil, err
		}
		for _, techTag := range techTags {
			roomTechTags = margeTechTagArray(roomTechTags, repository.TechTag{
				TechTagID: techTag.TechTagID.Int32,
				Language:  techTag.Language.String,
				Icon:      techTag.Icon.String,
			})
		}

		frameworks, err := store.ListAccountFrameworksByUserID(ctx, account.AccountID.String)
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

func parseRoomResponse(response domain.GetRoomResponse, room repository.Room, hackathon domain.RoomHackathonInfo) domain.GetRoomResponse {
	return domain.GetRoomResponse{
		RoomID:      room.RoomID,
		Title:       room.Title,
		Description: room.Description,
		MemberLimit: room.MemberLimit,
		IsDelete:    room.IsDelete,
		Hackathon:   hackathon,
		IsClosing:   room.IsClosing.Bool,
		NowMember:   response.NowMember,
	}
}

func getRoomMember(ctx context.Context, store transaction.Store, accountID string) (result []domain.NowRoomAccounts, err error) {
	accounts, err := store.GetRoomsAccountsByID(ctx, accountID)
	if err != nil {
		return
	}

	for _, account := range accounts {
		user, err := store.GetAccountsByID(ctx, account.AccountID.String)
		if err != nil {
			return nil, err
		}

		tags, err := parseTechTags(ctx, store, account.AccountID.String)
		if err != nil {
			return nil, err
		}

		frameworks, err := parseFrameworks(ctx, store, account.AccountID.String)
		if err != nil {
			return nil, err
		}

		result = append(result, domain.NowRoomAccounts{
			AccountID:  user.AccountID,
			Username:   user.Username,
			Icon:       user.Icon.String,
			IsOwner:    account.IsOwner,
			TechTags:   tags,
			Frameworks: frameworks,
		})
	}
	return
}

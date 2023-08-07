package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/domain"
	repository "github.com/hackhack-Geek-vol6/backend/gateways/repository/datasource"
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

func stackTagAndFrameworks(ctx context.Context, store transaction.Store, room repository.Room) ([]domain.RoomTechTags, []domain.RoomFramework, error) {
	var (
		roomTechTags   []domain.RoomTechTags
		roomFrameworks []domain.RoomFramework
	)
	accounts, err := store.GetRoomsAccountsByRoomID(ctx, room.RoomID)
	if err != nil {
		return nil, nil, err
	}

	for _, account := range accounts {
		techTags, err := store.ListAccountTagsByUserID(ctx, account.UserID.String)
		if err != nil {
			return nil, nil, err
		}
		for _, techTag := range techTags {
			roomTechTags = margeTechTagArray(roomTechTags, repository.TechTag{
				TechTagID: techTag.TechTagID.Int32,
				Language:  techTag.Language.String,
			})
		}

		frameworks, err := store.ListAccountFrameworksByUserID(ctx, account.UserID.String)
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

func margeRoomAccount(ctx context.Context, store transaction.Store, id uuid.UUID) (result []domain.NowRoomAccounts, err error) {
	nowMembers, err := store.GetRoomsAccountsByRoomID(ctx, id)
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

func (ru *roomUsecase) ListRooms(ctx context.Context, query domain.ListRoomsRequest) (result []domain.ListRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	rooms, err := ru.store.ListRoom(ctx, query.PageSize)
	if err != nil {
		return
	}

	for _, room := range rooms {
		techTags, frameworks, err := stackTagAndFrameworks(ctx, ru.store, room)
		if err != nil {
			return nil, err
		}

		hackathon, err := ru.store.GetHackathonByID(ctx, room.HackathonID)
		if err != nil {
			return nil, err
		}

		nowMember, err := margeRoomAccount(ctx, ru.store, room.RoomID)
		if err != nil {
			return nil, err
		}

		result = append(result, domain.ListRoomResponse{
			Rooms: domain.ListRoomRoomInfo{
				RoomID:      room.RoomID,
				Title:       room.Title,
				MemberLimit: room.MemberLimit,
				CreatedAt:   room.CreateAt,
			},
			Hackathon: domain.ListRoomHackathonInfo{
				HackathonID:   hackathon.HackathonID,
				HackathonName: hackathon.Name,
				Icon:          hackathon.Icon.String,
			},
			NowMember:         nowMember,
			MembersTechTags:   techTags,
			MembersFrameworks: frameworks,
		})
	}
	return
}

func (ru *roomUsecase) GetRoom(ctx context.Context, id uuid.UUID) (result domain.GetRoomResponse, err error) {
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

	tags, err := ru.store.GetHackathonStatusTagsByHackathonID(ctx, room.HackathonID)
	if err != nil {
		return
	}

	var statusTags []repository.StatusTag
	for _, tag := range tags {
		var statusTag repository.StatusTag
		statusTag, err = ru.store.GetStatusTagByStatusID(ctx, tag.StatusID)
		if err != nil {
			return domain.GetRoomResponse{}, err
		}

		statusTags = append(statusTags, statusTag)
	}

	techTags, frameworks, err := stackTagAndFrameworks(ctx, ru.store, room)
	if err != nil {
		return
	}

	nowMember, err := margeRoomAccount(ctx, ru.store, room.RoomID)
	if err != nil {
		return
	}
	result = domain.GetRoomResponse{
		RoomID:      room.RoomID,
		Title:       room.Title,
		Description: room.Description,
		MemberLimit: room.MemberLimit,
		IsStatus:    room.IsDelete,
		CreateAt:    room.CreateAt,
		Hackathon: domain.HackathonInfo{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Description: hackathon.Description,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			Tags:        statusTags,
		},
		NowMember:         nowMember,
		MembersTechTags:   techTags,
		MembersFrameworks: frameworks,
	}
	return
}

func (ru *roomUsecase) CreateRoom(ctx context.Context, body domain.CreateRoomRequestBody) (result domain.GetRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()
	return
}

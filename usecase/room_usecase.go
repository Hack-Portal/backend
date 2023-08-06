package usecase

import (
	"context"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type roomUsecase struct {
	store          db.Store
	contextTimeout time.Duration
}

func NewRoomUsercase(store db.Store, timeout time.Duration) domain.RoomUsecase {
	return &roomUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func stackTagAndFrameworks(ctx context.Context, store db.Store, room db.Rooms) ([]domain.RoomTechTags, []domain.RoomFramework, error) {
	var (
		roomTechTags   []db.RoomTechTags
		roomFrameworks []db.RoomFramework
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
			roomTechTags = margeTechTagArray(roomTechTags, db.TechTags{
				TechTagID: techTag.TechTagID.Int32,
				Language:  techTag.Language.String,
			})
		}

		frameworks, err := store.ListAccountFrameworksByUserID(ctx, account.UserID.String)
		if err != nil {
			return nil, nil, err
		}
		for _, framework := range frameworks {
			roomFrameworks = margeFrameworkArray(roomFrameworks, db.Frameworks{
				FrameworkID: framework.FrameworkID.Int32,
				TechTagID:   framework.TechTagID.Int32,
				Framework:   framework.Framework.String,
			})
		}
	}
	return roomTechTags, roomFrameworks, nil
}

func margeTechTagArray(roomTechTags []db.RoomTechTags, techtag db.TechTags) []db.RoomTechTags {
	for _, roomTechTag := range roomTechTags {
		if roomTechTag.TechTag == techtag {
			roomTechTag.Count++
		}
	}
	roomTechTags = append(roomTechTags, db.RoomTechTags{
		TechTag: techtag,
		Count:   1,
	})

	return roomTechTags
}

func margeFrameworkArray(roomFramework []db.RoomFramework, framework db.Frameworks) []db.RoomFramework {
	for _, roomFramework := range roomFramework {
		if roomFramework.Framework == framework {
			roomFramework.Count++
		}
	}
	roomFramework = append(roomFramework, db.RoomFramework{
		Framework: framework,
		Count:     1,
	})

	return roomFramework
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

		nowMember , err := ru.store.GetRoomsAccountsByRoomID(ctx,room.RoomID)
		if err != nil {
			return nil, err
		}
		
		result = append(result, domain.ListRoomResponse{
			domain.ListRoomRoomInfo{
				RoomID:      room.RoomID,
				Title:       room.Title,
				MemberLimit: room.MemberLimit,
				CreatedAt:   room.CreateAt,
			},
			domain.ListRoomHackathonInfo{
				HackathonID:   hackathon.HackathonID,
				HackathonName: hackathon.Name,
				Icon:          hackathon.Icon.String,
			},
			NowMember:
		})
	}

	return
}

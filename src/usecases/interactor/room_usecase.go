package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/domain/params"
	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/domain/response"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
)

type roomUsecase struct {
	store   transaction.Store
	l       logger.Logger
	timeout time.Duration
}

func NewRoomUsercase(store transaction.Store, l logger.Logger) inputport.RoomUsecase {
	return &roomUsecase{
		store:   store,
		l:       l,
		timeout: time.Duration(config.Config.Server.ContextTimeout),
	}
}

func (ru *roomUsecase) ListRooms(ctx context.Context, query request.ListRequest) ([]response.ListRoom, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	var result []response.ListRoom

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
		result = append(result, response.ListRoom{
			Rooms: response.ListRoomRoomInfo{
				RoomID:      room.RoomID,
				Title:       room.Title,
				MemberLimit: room.MemberLimit,
				IsClosing:   room.IsClosing.Bool,
				CreatedAt:   room.CreateAt,
			},
			Hackathon: response.ListRoomHackathonInfo{
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

func (ru *roomUsecase) GetRoom(ctx context.Context, id string) (result response.Room, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
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

	result = parseRoomResponse(result, room, response.RoomHackathonInfo{
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

func (ru *roomUsecase) CreateRoom(ctx context.Context, body params.CreateRoom) (result response.Room, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
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

	result = parseRoomResponse(result, room, response.RoomHackathonInfo{
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

func (ru *roomUsecase) UpdateRoom(ctx context.Context, body params.UpdateRoom) (result response.Room, err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
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

	result = parseRoomResponse(result, room, response.RoomHackathonInfo{
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

func (ru *roomUsecase) DeleteRoom(ctx context.Context, query params.DeleteRoom) error {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	return ru.store.DeleteRoomTx(ctx, query)
}

func (ru *roomUsecase) AddAccountInRoom(ctx context.Context, query params.AddAccountInRoom) error {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	return ru.store.AddAccountInRoom(ctx, query)
}

func (ru *roomUsecase) AddChat(ctx context.Context, body params.AddChat) error {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	data, err := ru.store.ReadDocsByRoomID(ctx, body.RoomID)
	if err != nil {
		return err
	}
	_, err = ru.store.CreateSubCollection(ctx, params.WriteFireStore{
		RoomID:  body.RoomID,
		Index:   data + 1,
		UID:     body.AccountID,
		Message: body.Message,
	})

	return err
}

func (ru *roomUsecase) DeleteRoomAccount(ctx context.Context, body params.DeleteRoomAccount) error {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	return ru.store.DeleteRoomsAccountsByID(ctx, repository.DeleteRoomsAccountsByIDParams{
		RoomID:    body.RoomID,
		AccountID: body.AccountID,
	})

}

func (ru *roomUsecase) CloseRoom(ctx context.Context, body params.CloseRoom) error {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	return ru.store.CloseRoom(ctx, body)
}

func (ru *roomUsecase) AddRoomAccountRole(ctx context.Context, body params.RoomAccountRole) error {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	return ru.store.AddRoomAccountRoleByID(ctx, body)
}

func (ru *roomUsecase) UpdateRoomAccountRole(ctx context.Context, body params.RoomAccountRole) (err error) {
	ctx, cancel := context.WithTimeout(ctx, ru.timeout)
	defer cancel()

	return ru.store.UpdateRoomsAccountRoleByID(ctx, body)
}

func parseRoles(ctx context.Context, store transaction.Store, id int32) (result []repository.Role, err error) {

	roles, err := store.ListRoomsAccountsRolesByID(ctx, id)
	if err != nil {
		return
	}
	for _, role := range roles {
		result = append(result, repository.Role{
			RoleID: role.RoleID,
			Role:   role.Role,
		})
	}
	return
}

func stackTagAndFrameworks(ctx context.Context, store transaction.Store, room repository.Room) ([]response.RoomTechTags, []response.RoomFramework, error) {
	var (
		roomTechTags   []response.RoomTechTags
		roomFrameworks []response.RoomFramework
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
				Icon:        framework.Icon.String,
			})
		}
	}
	return roomTechTags, roomFrameworks, nil
}

func margeTechTagArray(roomTechTags []response.RoomTechTags, techtag repository.TechTag) []response.RoomTechTags {
	for _, roomTechTag := range roomTechTags {
		if roomTechTag.TechTag == techtag {
			roomTechTag.Count++
		}
	}
	roomTechTags = append(roomTechTags, response.RoomTechTags{
		TechTag: techtag,
		Count:   1,
	})

	return roomTechTags
}

func margeFrameworkArray(roomFramework []response.RoomFramework, framework repository.Framework) []response.RoomFramework {
	for _, roomFramework := range roomFramework {
		if roomFramework.Framework == framework {
			roomFramework.Count++
		}
	}
	roomFramework = append(roomFramework, response.RoomFramework{
		Framework: framework,
		Count:     1,
	})

	return roomFramework
}

func parseRoomResponse(resp response.Room, room repository.Room, hackathon response.RoomHackathonInfo) response.Room {
	return response.Room{
		RoomID:      room.RoomID,
		Title:       room.Title,
		Description: room.Description,
		MemberLimit: room.MemberLimit,
		IsDelete:    room.IsDelete,
		Hackathon:   hackathon,
		IsClosing:   room.IsClosing.Bool,
		NowMember:   resp.NowMember,
	}
}

func getRoomMember(ctx context.Context, store transaction.Store, accountID string) (result []response.NowRoomAccounts, err error) {
	accounts, err := store.GetRoomsAccountsByID(ctx, accountID)
	if err != nil {
		return
	}

	for _, account := range accounts {
		user, err := store.GetAccountsByID(ctx, account.AccountID.String)
		if err != nil {
			return nil, err
		}

		roles, err := parseRoles(ctx, store, account.RoomsAccountID)
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

		result = append(result, response.NowRoomAccounts{
			AccountID:  user.AccountID,
			Username:   user.Username,
			Icon:       user.Icon.String,
			IsOwner:    account.IsOwner,
			Roles:      roles,
			TechTags:   tags,
			Frameworks: frameworks,
		})
	}
	return
}

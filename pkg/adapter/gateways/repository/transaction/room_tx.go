package transaction

import (
	"context"
	"errors"

	"github.com/google/uuid"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

func stackTagAndFrameworks(ctx context.Context, q *repository.Queries, room repository.Room) ([]domain.RoomTechTags, []domain.RoomFramework, error) {
	var (
		roomTechTags   []domain.RoomTechTags
		roomFrameworks []domain.RoomFramework
	)
	accounts, err := q.GetRoomsAccountsByRoomID(ctx, room.RoomID)
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
	nowMembers, err := q.GetRoomsAccountsByRoomID(ctx, id)
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

func getHackathonTag(ctx context.Context, q *repository.Queries, id int32) (result []repository.StatusTag, err error) {
	tags, err := q.GetHackathonStatusTagsByHackathonID(ctx, id)
	if err != nil {
		return
	}

	for _, tag := range tags {
		statusTag, err := q.GetStatusTagByStatusID(ctx, tag.StatusID)
		if err != nil {
			return nil, err
		}
		result = append(result, statusTag)
	}
	return
}

func getRoomMember(ctx context.Context, q *repository.Queries, id uuid.UUID) (result []domain.NowRoomAccounts, err error) {
	accounts, err := q.GetRoomsAccountsByRoomID(ctx, id)
	if err != nil {
		return
	}

	for _, account := range accounts {
		user, err := q.GetAccountByID(ctx, account.UserID.String)
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

func compRoom(request domain.UpdateRoomParam, latest repository.Room, members int32) (result repository.UpdateRoomByIDParams, err error) {
	result.RoomID = latest.RoomID

	if len(request.Title) != 0 {
		if latest.Title != request.Title {
			result.Title = request.Title
		}
	} else {
		result.Title = latest.Title
	}

	if len(request.Description) != 0 {
		if latest.Description != request.Description {
			result.Description = request.Description
		}
	} else {
		result.Description = latest.Description
	}

	if request.MemberLimit != 0 {
		if request.MemberLimit > int32(members) {
			result.MemberLimit = request.MemberLimit
		} else {
			err = errors.New("現在の加入メンバーを下回る変更はできない")
			return
		}
	} else {
		result.MemberLimit = latest.MemberLimit
	}

	if request.HackathonID != 0 {
		if latest.HackathonID != request.HackathonID {
			result.HackathonID = request.HackathonID
		}
	} else {
		result.HackathonID = latest.HackathonID
	}

	return
}

func checkOwner(members []domain.NowRoomAccounts, id string) bool {
	for _, member := range members {
		if member.UserID == id {
			return member.IsOwner
		}
	}
	return false
}

func checkDuplication(members []domain.NowRoomAccounts, id string) bool {
	for _, member := range members {
		if member.UserID == id {
			return true
		}
	}
	return false
}

func (store *SQLStore) CreateRoomTx(ctx context.Context, args domain.CreateRoomParam) (domain.GetRoomResponse, error) {
	var result domain.GetRoomResponse
	err := store.execTx(ctx, func(q *repository.Queries) error {

		room, err := q.CreateRoom(ctx, repository.CreateRoomParams{
			RoomID:      args.RoomID,
			HackathonID: args.HackathonID,
			Title:       args.Title,
			Description: args.Description,
			MemberLimit: args.MemberLimit,
			IsDelete:    false,
		})
		if err != nil {
			return err
		}

		hackathon, err := q.GetHackathonByID(ctx, args.HackathonID)
		if err != nil {
			return err
		}

		hackathonTag, err := getHackathonTag(ctx, q, args.HackathonID)
		if err != nil {
			return err
		}

		_, err = q.CreateRoomsAccounts(ctx, repository.CreateRoomsAccountsParams{
			UserID:  args.OwnerID,
			RoomID:  room.RoomID,
			IsOwner: true,
		})

		if err != nil {
			return err
		}

		techTags, frameworks, err := stackTagAndFrameworks(ctx, q, room)
		if err != nil {
			return err
		}

		members, err := getRoomMember(ctx, q, room.RoomID)
		if err != nil {
			return err
		}

		result = parseRoomResponse(result, room, domain.HackathonInfo{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Link:        hackathon.Link,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			Tags:        hackathonTag,
		})
		result.NowMember = members
		result.MembersTechTags = techTags
		result.MembersFrameworks = frameworks

		return nil
	})
	return result, err
}

func (store *SQLStore) GetRoomTx(ctx context.Context, id uuid.UUID) (domain.GetRoomResponse, error) {
	var result domain.GetRoomResponse
	err := store.execTx(ctx, func(q *repository.Queries) error {

		room, err := q.GetRoomsByID(ctx, id)
		if err != nil {
			return err
		}

		hackathon, err := q.GetHackathonByID(ctx, room.HackathonID)
		if err != nil {
			return err
		}

		hackathonTag, err := getHackathonTag(ctx, q, room.HackathonID)
		if err != nil {
			return err
		}

		techTags, frameworks, err := stackTagAndFrameworks(ctx, q, room)
		if err != nil {
			return err
		}

		members, err := getRoomMember(ctx, q, room.RoomID)
		if err != nil {
			return err
		}

		result = parseRoomResponse(result, room, domain.HackathonInfo{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Link:        hackathon.Link,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			Tags:        hackathonTag,
		})
		result.NowMember = members
		result.MembersTechTags = techTags
		result.MembersFrameworks = frameworks

		return nil
	})
	return result, err
}

func (store *SQLStore) ListRoomTx(ctx context.Context, query domain.ListRoomsRequest) ([]domain.ListRoomResponse, error) {
	var result []domain.ListRoomResponse
	err := store.execTx(ctx, func(q *repository.Queries) error {

		rooms, err := q.ListRoom(ctx, query.PageSize)
		if err != nil {
			return err
		}

		for _, room := range rooms {
			hackathon, err := q.GetHackathonByID(ctx, room.HackathonID)
			if err != nil {
				return err
			}

			techTags, frameworks, err := stackTagAndFrameworks(ctx, q, room)
			if err != nil {
				return err
			}

			members, err := getRoomMember(ctx, q, room.RoomID)
			if err != nil {
				return err
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
				NowMember:         members,
				MembersTechTags:   techTags,
				MembersFrameworks: frameworks,
			})
		}
		return nil
	})
	return result, err
}

func (store *SQLStore) UpdateRoomTx(ctx context.Context, body domain.UpdateRoomParam) (domain.GetRoomResponse, error) {
	var result domain.GetRoomResponse
	err := store.execTx(ctx, func(q *repository.Queries) error {

		latest, err := q.GetRoomsByID(ctx, body.RoomID)
		if err != nil {
			return err
		}

		members, err := getRoomMember(ctx, q, latest.RoomID)
		if err != nil {
			return err
		}

		owner, err := q.GetAccountByEmail(ctx, body.OwnerEmail)
		if err != nil {
			return err
		}

		if !checkOwner(members, owner.UserID) {
			err := errors.New("あんたオーナーとちゃうやん")
			return err
		}

		args, err := compRoom(body, latest, int32(len(members)))
		if err != nil {
			return err
		}

		room, err := q.UpdateRoomByID(ctx, args)
		if err != nil {
			return err
		}

		hackathon, err := q.GetHackathonByID(ctx, room.HackathonID)
		if err != nil {
			return err
		}

		hackathonTag, err := getHackathonTag(ctx, q, room.HackathonID)
		if err != nil {
			return err
		}

		techTags, frameworks, err := stackTagAndFrameworks(ctx, q, room)
		if err != nil {
			return err
		}

		result = parseRoomResponse(result, room, domain.HackathonInfo{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Link:        hackathon.Link,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			Tags:        hackathonTag,
		})
		result.NowMember = members
		result.MembersTechTags = techTags
		result.MembersFrameworks = frameworks
		return nil
	})
	return result, err
}

func (store *SQLStore) DeleteRoomTx(ctx context.Context, args domain.DeleteRoomParam) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {

		owner, err := q.GetAccountByEmail(ctx, args.OwnerEmail)
		if err != nil {
			return err
		}

		members, err := getRoomMember(ctx, q, args.RoomID)
		if err != nil {
			return err
		}

		if !checkOwner(members, owner.UserID) {
			err := errors.New("あんたオーナーとちゃうやん")
			return err
		}

		_, err = q.SoftDeleteRoomByID(ctx, args.RoomID)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (store *SQLStore) AddAccountInRoom(ctx context.Context, args domain.AddAccountInRoomParam) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {

		members, err := getRoomMember(ctx, q, args.RoomID)
		if err != nil {
			return err
		}

		if !checkDuplication(members, args.UserID) {
			err := errors.New("あんたすでにルームおるやん")
			return err
		}

		_, err = q.CreateRoomsAccounts(ctx, repository.CreateRoomsAccountsParams{
			UserID:  args.UserID,
			RoomID:  args.RoomID,
			IsOwner: false,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

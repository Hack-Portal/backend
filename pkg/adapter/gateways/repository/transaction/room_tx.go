package transaction

import (
	"context"
	"errors"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

func compRoom(request domain.UpdateRoomParam, latest repository.Room, members int32) (result repository.UpdateRoomsByIDParams, err error) {
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

func checkOwner(members []repository.GetRoomsAccountsByIDRow, id string) bool {
	for _, member := range members {
		if member.UserID.String == id {
			return member.IsOwner
		}
	}
	return false
}

func checkDuplication(members []repository.GetRoomsAccountsByIDRow, id string) bool {
	for _, member := range members {
		if member.UserID.String == id {
			return true
		}
	}
	return false
}

func (store *SQLStore) CreateRoomTx(ctx context.Context, args domain.CreateRoomParam) (domain.GetRoomResponse, error) {
	var result domain.GetRoomResponse
	err := store.execTx(ctx, func(q *repository.Queries) error {

		room, err := q.CreateRooms(ctx, repository.CreateRoomsParams{
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

		_, err = q.CreateRoomsAccounts(ctx, repository.CreateRoomsAccountsParams{
			UserID:  args.OwnerID,
			RoomID:  room.RoomID,
			IsOwner: true,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

func (store *SQLStore) UpdateRoomTx(ctx context.Context, body domain.UpdateRoomParam) (repository.Room, error) {
	var room repository.Room
	err := store.execTx(ctx, func(q *repository.Queries) error {

		latest, err := q.GetRoomsByID(ctx, body.RoomID)
		if err != nil {
			return err
		}

		members, err := q.GetRoomsAccountsByID(ctx, latest.RoomID)
		if err != nil {
			return err
		}

		owner, err := q.GetAccountsByEmail(ctx, body.OwnerEmail)
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

		room, err = q.UpdateRoomsByID(ctx, args)
		if err != nil {
			return err
		}
		return nil
	})
	return room, err
}

func (store *SQLStore) DeleteRoomTx(ctx context.Context, args domain.DeleteRoomParam) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {

		owner, err := q.GetAccountsByEmail(ctx, args.OwnerEmail)
		if err != nil {
			return err
		}

		members, err := q.GetRoomsAccountsByID(ctx, args.RoomID)
		if err != nil {
			return err
		}

		if !checkOwner(members, owner.UserID) {
			err := errors.New("あんたオーナーとちゃうやん")
			return err
		}

		_, err = q.DeleteRoomsByID(ctx, args.RoomID)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (store *SQLStore) AddAccountInRoom(ctx context.Context, args domain.AddAccountInRoomParam) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {

		members, err := q.GetRoomsAccountsByID(ctx, args.RoomID)
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

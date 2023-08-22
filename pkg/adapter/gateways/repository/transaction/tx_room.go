package transaction

import (
	"context"
	"errors"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
)

func compRoom(request domain.UpdateRoomParam, latest repository.Room, members int32) (result repository.UpdateRoomsByIDParams, err error) {
	result = repository.UpdateRoomsByIDParams{
		HackathonID: latest.HackathonID,
		Title:       latest.Title,
		Description: latest.Description,
		MemberLimit: latest.MemberLimit,
		RoomID:      request.RoomID,
		UpdateAt:    time.Now(),
	}

	if util.CheckDiff(latest.Title, request.Title) {
		result.Title = request.Title
	}

	if util.CheckDiff(latest.Description, request.Description) {
		result.Description = request.Description
	}

	if request.MemberLimit != 0 {
		if request.MemberLimit > int32(members) {
			result.MemberLimit = request.MemberLimit
		} else {
			err = errors.New("現在の加入メンバーを下回る変更はできない")
			return
		}
	}

	if request.HackathonID != 0 {
		if latest.HackathonID != request.HackathonID {
			result.HackathonID = request.HackathonID
		}
	}

	return
}

func checkOwner(members []repository.GetRoomsAccountsByIDRow, id string) bool {
	for _, member := range members {
		if member.AccountID.String == id {
			return member.IsOwner
		}
	}
	return false
}

func checkDuplication(members []repository.GetRoomsAccountsByIDRow, id string) bool {
	for _, member := range members {
		if member.AccountID.String == id {
			return true
		}
	}
	return false
}

func (store *SQLStore) CreateRoomTx(ctx context.Context, args domain.CreateRoomParam) (repository.Room, error) {
	var room repository.Room
	err := store.execTx(ctx, func(q *repository.Queries) error {
		var err error
		room, err = q.CreateRooms(ctx, repository.CreateRoomsParams{
			RoomID:      args.RoomID,
			HackathonID: args.HackathonID,
			Title:       args.Title,
			Description: args.Description,
			MemberLimit: args.MemberLimit,
			IncludeRate: args.IncludeRate,
		})
		if err != nil {
			return err
		}

		_, err = q.CreateRoomsAccounts(ctx, repository.CreateRoomsAccountsParams{
			AccountID: args.OwnerID,
			RoomID:    room.RoomID,
			IsOwner:   true,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return room, err
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
		if !checkOwner(members, owner.AccountID) {
			return errors.New("あんたオーナーとちゃうやん")
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

		if !checkOwner(members, owner.AccountID) {
			return errors.New("あんたオーナーとちゃうやん")
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

		if checkDuplication(members, args.AccountID) {
			return errors.New("あんたすでにルームおるやん")
		}

		_, err = q.CreateRoomsAccounts(ctx, repository.CreateRoomsAccountsParams{
			AccountID: args.AccountID,
			RoomID:    args.RoomID,
			IsOwner:   false,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

package transaction

import (
	"context"
	"database/sql"
	"errors"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
)

func compRoom(request params.UpdateRoom, latest repository.Room) (result repository.UpdateRoomsByIDParams, err error) {
	result = repository.UpdateRoomsByIDParams{
		HackathonID: latest.HackathonID,
		Title:       latest.Title,
		Description: latest.Description,
		MemberLimit: latest.MemberLimit,
		RoomID:      request.RoomID,
		IsClosing:   sql.NullBool{Bool: request.IsClosing, Valid: true},
		UpdateAt:    time.Now(),
	}

	if util.CheckDiff(latest.Title, request.Title) {
		result.Title = request.Title
	}

	if util.CheckDiff(latest.Description, request.Description) {
		result.Description = request.Description
	}

	if request.MemberLimit != 0 {
		result.MemberLimit = request.MemberLimit
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

func (store *SQLStore) CreateRoomTx(ctx context.Context, args params.CreateRoom) (repository.Room, error) {
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
			IsClosing:   sql.NullBool{Bool: false, Valid: true},
		})
		if err != nil {
			return err
		}

		if _, err = q.CreateRoomsAccounts(ctx, repository.CreateRoomsAccountsParams{
			AccountID: args.OwnerID,
			RoomID:    room.RoomID,
			IsOwner:   true,
		}); err != nil {
			return err
		}

		return nil
	})
	return room, err
}

func (store *SQLStore) UpdateRoomTx(ctx context.Context, body params.UpdateRoom) (repository.Room, error) {
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

		args, err := compRoom(body, latest)
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

func (store *SQLStore) DeleteRoomTx(ctx context.Context, args params.DeleteRoom) error {
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

func (store *SQLStore) AddAccountInRoom(ctx context.Context, args params.AddAccountInRoom) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {

		rooms, err := q.GetRoomsByID(ctx, args.RoomID)
		if err != nil {
			return err
		}

		if rooms.IsClosing.Bool {
			return errors.New("すまんが、もう閉め切ってんねんなこのルーム")
		}

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

func (store *SQLStore) CloseRoom(ctx context.Context, args params.CloseRoom) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {

		rooms, err := q.GetRoomsByID(ctx, args.RoomID)
		if err != nil {
			return err
		}

		if rooms.IsClosing.Bool {
			return errors.New("すまんが、もう閉め切ってんねんなこのルーム")
		}

		members, err := q.GetRoomsAccountsByID(ctx, args.RoomID)
		if err != nil {
			return err
		}

		if len(args.AccountID) > len(members) {
			return errors.New("募集人数より多ないか？")
		}
		for _, member := range members {
			for _, accountID := range args.AccountID {
				if member.AccountID.String != accountID {
					if err := q.DeleteRoomsAccountsByID(ctx, repository.DeleteRoomsAccountsByIDParams{
						RoomID:    args.RoomID,
						AccountID: member.AccountID.String,
					}); err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
	return err
}

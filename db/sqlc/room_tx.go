package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateRoomTxParams struct {
	// ルーム登録部分
	Rooms
	// RoomsAccounts登録部分
	UserID string
}
type RoomTechTags struct {
	TechTag TechTags `json:"tech_tag"`
	Count   int32    `json:"count"`
}
type RoomFramework struct {
	Framework Frameworks `json:"framework"`
	Count     int32      `json:"count"`
}
type RoomHackathonData struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CraeteRoomTxResult struct {
	Rooms
	Hackathon       RoomHackathonData
	RoomsAccounts   []GetRoomsAccountsByRoomIDRow
	RoomsTechTags   []RoomTechTags
	RoomsFrameworks []RoomFramework
}

// TechTagの配列にマージする
func margeTechTagArray(roomTechTags []RoomTechTags, techtag TechTags) []RoomTechTags {
	for _, roomTechTag := range roomTechTags {
		if roomTechTag.TechTag == techtag {
			roomTechTag.Count++
		}
	}
	roomTechTags = append(roomTechTags, RoomTechTags{
		TechTag: techtag,
		Count:   1,
	})

	return roomTechTags
}

// フレームワークの配列にマージする
func margeFrameworkArray(roomFramework []RoomFramework, framework Frameworks) []RoomFramework {
	for _, roomFramework := range roomFramework {
		if roomFramework.Framework == framework {
			roomFramework.Count++
		}
	}
	roomFramework = append(roomFramework, RoomFramework{
		Framework: framework,
		Count:     1,
	})

	return roomFramework
}

func (store *SQLStore) CreateRoomTx(ctx context.Context, arg CreateRoomTxParams) (CraeteRoomTxResult, error) {
	var result CraeteRoomTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// ハッカソンデータを送る
		hackathon, err := q.GetHackathonByID(ctx, arg.HackathonID)
		if err != nil {
			return err
		}
		result.Hackathon = RoomHackathonData{
			ID:   hackathon.HackathonID,
			Name: hackathon.Name,
			Icon: hackathon.Icon.String,
		}

		// ルームを登録する
		result.Rooms, err = q.CreateRoom(ctx, CreateRoomParams{
			RoomID:      arg.RoomID,
			HackathonID: hackathon.HackathonID,
			Title:       arg.Title,
			Description: arg.Description,
			MemberLimit: arg.MemberLimit,
			IsStatus:    true,
		})
		if err != nil {
			return err
		}

		// ルームのオーナーを登録する
		_, err = q.CreateRoomsAccounts(ctx, CreateRoomsAccountsParams{
			UserID:  arg.UserID,
			RoomID:  result.RoomID,
			IsOwner: true,
		})

		if err != nil {
			return err
		}

		result.RoomsAccounts, err = q.GetRoomsAccountsByRoomID(ctx, result.RoomID)
		if err != nil {
			return err
		}

		// ルーム内のユーザをもとにユーザの持つ技術タグとフレームワークタグを配列に落とし込む（力業
		for _, account := range result.RoomsAccounts {
			techTags, err := q.ListAccountTagsByUserID(ctx, account.UserID.String)
			if err != nil {
				return err
			}
			for _, techTag := range techTags {
				result.RoomsTechTags = margeTechTagArray(result.RoomsTechTags, TechTags{
					TechTagID: techTag.TechTagID.Int32,
					Language:  techTag.Language.String,
				})
			}

			frameworks, err := q.ListAccountFrameworksByUserID(ctx, account.UserID.String)
			if err != nil {
				return err
			}
			for _, framework := range frameworks {
				result.RoomsFrameworks = margeFrameworkArray(result.RoomsFrameworks, Frameworks{
					FrameworkID: framework.FrameworkID.Int32,
					TechTagID:   framework.TechTagID.Int32,
					Framework:   framework.Framework.String,
				})
			}
		}
		return nil
	})
	return result, err
}

type ListRoomTxParam struct {
}

type ListRoomTxRoomInfo struct {
	RoomID      uuid.UUID `json:"room_id"`
	Title       string    `josn:"title"`
	MemberLimit int32     `json:"member_limit"`
	CreatedAt   time.Time `json:"created_at"`
}
type ListRoomTxHacathonInfo struct {
	HackathonID   int32  `json:"hackathon_id"`
	HackathonName string `json:"hackathon_name"`
	Icon          string `json:"icon"`
}
type ListRoomTxResult struct {
	Rooms             ListRoomTxRoomInfo            `json:"rooms"`
	Hackathon         ListRoomTxHacathonInfo        `json:"hackathon"`
	NowMember         []GetRoomsAccountsByRoomIDRow `json:"now_member"`
	MembersTechTags   []RoomTechTags                `json:"members_tech_tags"`
	MembersFrameworks []RoomFramework               `json:"members_frameworks"`
}

func (store *SQLStore) ListRoomTx(ctx context.Context, arg ListRoomTxParam) ([]ListRoomTxResult, error) {
	var result []ListRoomTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// ルーム一覧を取得してくる
		rooms, err := q.ListRoom(ctx, 100)
		if err != nil {
			return err
		}
		// それぞれのルームの確認
		for _, room := range rooms {
			var oneRoomInfos ListRoomTxResult
			oneRoomInfos.Rooms = ListRoomTxRoomInfo{
				RoomID:      room.RoomID,
				Title:       room.Title,
				MemberLimit: room.MemberLimit,
				CreatedAt:   room.CreateAt,
			}
			hackathon, err := q.GetHackathonByID(ctx, room.HackathonID)
			if err != nil {
				return err
			}
			// ハッカソンの追加
			oneRoomInfos.Hackathon = ListRoomTxHacathonInfo{
				HackathonID:   hackathon.HackathonID,
				HackathonName: hackathon.Name,
				Icon:          hackathon.Icon.String,
			}

			members, err := q.GetRoomsAccountsByRoomID(ctx, room.RoomID)
			if err != nil {
				return err
			}
			// アカウントの追加
			for _, account := range members {
				// タグの追加
				techTags, err := q.ListAccountTagsByUserID(ctx, account.UserID.String)
				if err != nil {
					return err
				}
				for _, techTag := range techTags {
					oneRoomInfos.MembersTechTags = margeTechTagArray(oneRoomInfos.MembersTechTags, TechTags{
						TechTagID: techTag.TechTagID.Int32,
						Language:  techTag.Language.String,
					})
				}
				// FWの追加
				frameworks, err := q.ListAccountFrameworksByUserID(ctx, account.UserID.String)
				if err != nil {
					return err
				}
				for _, framework := range frameworks {
					oneRoomInfos.MembersFrameworks = margeFrameworkArray(oneRoomInfos.MembersFrameworks, Frameworks{
						FrameworkID: framework.FrameworkID.Int32,
						TechTagID:   framework.TechTagID.Int32,
						Framework:   framework.Framework.String,
					})
				}
			}
			result = append(result, oneRoomInfos)
		}
		return err
	})
	return result, err
}

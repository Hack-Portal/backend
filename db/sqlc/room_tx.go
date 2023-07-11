package db

import "context"

type CreateRoomTxParams struct {
	// ルーム登録部分
	Rooms
	// RoomsAccounts登録部分
	UserID string
	// テックタグ登録部分
	RoomsTechTags []int32
	// フレームワーク登録部分
	RoomsFrameworks []int32
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
	Icon []byte `json:"icon"`
}

type CraeteRoomTxResult struct {
	Rooms
	Hackathon       RoomHackathonData
	RoomsAccounts   []GetRoomsAccountsRow
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
		hackathon, err := q.GetHackathon(ctx, arg.HackathonID)
		if err != nil {
			return err
		}
		result.Hackathon = RoomHackathonData{
			ID:   hackathon.HackathonID,
			Name: hackathon.Name,
			Icon: hackathon.Icon,
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
		result.RoomsAccounts, err = q.GetRoomsAccounts(ctx, result.RoomID)

		if err != nil {
			return err
		}

		// ルームＩＤからテックタグのレコードを登録する
		for _, techtag := range arg.RoomsTechTags {
			accountTag, err := q.CreateRoomsTechTag(ctx, CreateRoomsTechTagParams{
				RoomID:    result.RoomID,
				TechTagID: techtag,
			})
			if err != nil {
				return err
			}
			techtag, err := q.GetTechTag(ctx, accountTag.TechTagID)
			if err != nil {
				return err
			}

			result.RoomsTechTags = margeTechTagArray(result.RoomsTechTags, techtag)
		}

		// ルームＩＤからフレームワークのレコードを登録する
		for _, accountFrameworkTag := range arg.RoomsFrameworks {
			accountFramework, err := q.CreateRoomsFramework(ctx, CreateRoomsFrameworkParams{
				RoomID:      result.RoomID,
				FrameworkID: accountFrameworkTag,
			})
			if err != nil {
				return err
			}
			framework, err := q.GetFrameworks(ctx, accountFramework.FrameworkID)
			if err != nil {
				return err
			}
			result.RoomsFrameworks = margeFrameworkArray(result.RoomsFrameworks, framework)
		}

		return nil
	})
	return result, err
}

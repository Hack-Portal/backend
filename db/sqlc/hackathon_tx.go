package db

import "context"

type CreateHackathonTxParams struct {
	// ハッカソン登録部分
	Hackathons
	// status_tag登録用
	HackathonStatusTag []int32
}

type CreateHackathonTxResult struct {
	Hackathons
	HackathonStatusTag []StatusTags
}

// ハッカソン登録時のトランザクション
func (store *SQLStore) CreateHackathonTx(ctx context.Context, arg CreateHackathonTxParams) (CreateHackathonTxResult, error) {
	var result CreateHackathonTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// ハッカソンを登録する
		result.Hackathons, err = q.CreateHackathon(ctx, CreateHackathonParams{
			Name:        arg.Name,
			Icon:        arg.Icon,
			Description: arg.Description,
			Link:        arg.Link,
			Expired:     arg.Expired,
			StartDate:   arg.StartDate,
			Term:        arg.Term,
		})
		if err != nil {
			return err
		}
		// ハッカソンIDからステータスタグのレコードを登録する
		for _, status_id := range arg.HackathonStatusTag {
			_, err = q.CreateHackathonStatusTag(ctx, CreateHackathonStatusTagParams{
				HackathonID: result.HackathonID,
				StatusID:    status_id,
			})
			if err != nil {
				return err
			}
		}
		statusTag, err := q.GetStatusTagsByHackathonID(ctx, result.HackathonID)
		if err != nil {
			return err
		}
		result.HackathonStatusTag = statusTag
		return nil
	})
	return result, err
}

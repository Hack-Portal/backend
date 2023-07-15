package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/hackhack-Geek-vol6/backend/util"
)

type CreateHackathonTxParams struct {
	// ハッカソン登録部分
	Name        string    `json:"name"`
	Icon        []byte    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`

	// status_tag登録用
	HackathonStatusTag []int32
}

type CreateHackathonTxResult struct {
	Hackathons
	HackathonStatusTag []StatusTags
}

// ハッカソン登録時のトランザクション
func (store *SQLStore) CreateHackathonTx(ctx context.Context, config *util.EnvConfig, arg CreateHackathonTxParams) (CreateHackathonTxResult, error) {
	var result CreateHackathonTxResult
	id, err := uuid.NewV1()
	if err != nil {
		return result, err
	}
	hackathonToken, err := store.UploadImage(ctx, arg.Icon, id.String()+".jpg")
	if err != nil {
		return result, err
	}
	err = store.execTx(ctx, func(q *Queries) error {
		var err error
		// ハッカソンを登録する
		result.Hackathons, err = q.CreateHackathon(ctx, CreateHackathonParams{
			Name: arg.Name,
			Icon: sql.NullString{
				String: fmt.Sprintf("%s/%s.jpg?alt=media&token=%s", config.BasePath, id, hackathonToken),
				Valid:  true,
			},
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

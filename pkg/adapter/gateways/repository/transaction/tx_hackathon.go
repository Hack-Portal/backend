package transaction

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
)

func (store *SQLStore) CreateHackathonTx(ctx context.Context, args params.CreateHackathon) (repository.Hackathon, error) {
	var hackathon repository.Hackathon
	err := store.execTx(ctx, func(q *repository.Queries) error {
		var err error
		hackathon, err = q.CreateHackathons(ctx, args.Hackathon)
		if err != nil {
			return err
		}

		for _, tag := range args.StatusTags {
			_, err = q.CreateHackathonStatusTags(ctx, repository.CreateHackathonStatusTagsParams{
				HackathonID: hackathon.HackathonID,
				StatusID:    tag,
			})

			if err != nil {
				return err
			}
		}
		return nil
	})
	return hackathon, err
}

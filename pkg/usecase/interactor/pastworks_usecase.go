package usecase

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

func parsePastworkMembers(ctx context.Context, q *repository.Queries, id []string) (result []domain.PastWorkMembers, err error) {
	for _, userid := range id {
		account, err := q.GetAccountsByID(ctx, userid)
		if err != nil {
			return nil, err
		}
		result = append(result, domain.PastWorkMembers{AccountID: account.AccountID, Name: account.Username, Icon: account.Icon.String})
	}
	return
}

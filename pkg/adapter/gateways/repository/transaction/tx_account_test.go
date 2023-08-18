package transaction

import (
	"context"
	"fmt"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountTx(t *testing.T) {
	testCases := []struct {
		name          string
		arg           domain.CreateAccountParams
		checkReturens func(t *testing.T, arg domain.CreateAccountParams, returns repository.Account)
	}{
		{
			name: "success",
			arg: domain.CreateAccountParams{
				AccountInfo: repository.CreateAccountsParams{
					AccountID:       util.RandomString(10),
					Email:           util.RandomEmail(),
					Username:        util.RandomString(10),
					Icon:            dbutil.ToSqlNullString(util.RandomString(10)),
					ExplanatoryText: dbutil.ToSqlNullString(util.RandomString(10)),
					LocateID:        int32(util.Random(47)),
					Rate:            0,
					Character:       dbutil.ToSqlNullInt32(10),
					ShowLocate:      false,
					ShowRate:        false,
				},
				AccountTechTag:      util.RandomSelection(14, 5),
				AccountFrameworkTag: util.RandomSelection(52, 5),
			},
			checkReturens: func(t *testing.T, arg domain.CreateAccountParams, returns repository.Account) {

			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			account, err := TestStore.CreateAccountTx(context.Background(), tc.arg)
			require.NoError(t, err)
			require.NotEmpty(t, account)

			fmt.Println(account)

			tc.checkReturens(t, tc.arg, account)
		})
	}
}

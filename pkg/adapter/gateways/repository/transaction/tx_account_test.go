package transaction

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func randomAccount(t *testing.T) (domain.CreateAccountParams, repository.Account) {
	arg := domain.CreateAccountParams{
		AccountInfo: repository.CreateAccountsParams{
			AccountID:       util.RandomString(10),
			Email:           util.RandomEmail(),
			Username:        util.RandomString(10),
			Icon:            dbutil.ToSqlNullString(util.RandomString(10)),
			ExplanatoryText: dbutil.ToSqlNullString(util.RandomString(10)),
			LocateID:        int32(util.Random(47)),
			Rate:            0,
			Character:       dbutil.ToSqlNullInt32(int32(util.Random(5) + 1)),
			ShowLocate:      false,
			ShowRate:        false,
		},
		AccountTechTag:      util.RandomSelection(14, 5),
		AccountFrameworkTag: util.RandomSelection(52, 5),
	}

	account, err := store.CreateAccountTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	return arg, account
}
func TestCreateAccountTx(t *testing.T) {
	arg, account := randomAccount(t)

	tags, err := store.ListAccountTagsByUserID(context.Background(), account.AccountID)
	require.NoError(t, err)

	frameworks, err := store.ListAccountFrameworksByUserID(context.Background(), account.AccountID)
	require.NoError(t, err)

	require.Equal(t, arg.AccountInfo.AccountID, account.AccountID)
	require.Equal(t, arg.AccountInfo.Username, account.Username)
	require.Equal(t, arg.AccountInfo.Icon, account.Icon)
	require.Equal(t, arg.AccountInfo.ExplanatoryText, account.ExplanatoryText)
	require.Equal(t, arg.AccountInfo.LocateID, account.LocateID)
	require.Equal(t, arg.AccountInfo.Rate, account.Rate)
	require.Equal(t, arg.AccountInfo.Character, account.Character)
	require.Equal(t, arg.AccountInfo.ShowLocate, account.ShowLocate)
	require.Equal(t, arg.AccountInfo.ShowRate, account.ShowRate)
	require.Equal(t, arg.AccountInfo.ShowLocate, account.ShowLocate)
	require.NotEmpty(t, account.CreateAt)
	require.NotEmpty(t, account.UpdateAt)
	require.Equal(t, false, account.IsDelete)

	require.Len(t, tags, len(arg.AccountTechTag))
	require.Len(t, frameworks, len(arg.AccountFrameworkTag))
}

func TestUpdateAccountTx(t *testing.T) {
	_, account := randomAccount(t)

	arg := domain.UpdateAccountParam{
		AccountInfo: repository.Account{
			AccountID:       account.AccountID,
			Email:           account.Email,
			Username:        util.RandomString(10),
			Icon:            dbutil.ToSqlNullString(util.RandomString(10)),
			ExplanatoryText: dbutil.ToSqlNullString(util.RandomString(10)),
			LocateID:        int32(util.Random(47)),
			Rate:            account.Rate,
			Character:       dbutil.ToSqlNullInt32(int32(util.Random(5) + 1)),
			ShowLocate:      false,
			ShowRate:        false,
		},
		AccountTechTag:      util.RandomSelection(14, 5),
		AccountFrameworkTag: util.RandomSelection(52, 5),
	}

	account, err := store.UpdateAccountTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	tags, err := store.ListAccountTagsByUserID(context.Background(), account.AccountID)
	require.NoError(t, err)

	frameworks, err := store.ListAccountFrameworksByUserID(context.Background(), account.AccountID)
	require.NoError(t, err)

	require.Equal(t, arg.AccountInfo.AccountID, account.AccountID)
	require.Equal(t, arg.AccountInfo.Username, account.Username)
	require.Equal(t, arg.AccountInfo.Icon, account.Icon)
	require.Equal(t, arg.AccountInfo.ExplanatoryText, account.ExplanatoryText)
	require.Equal(t, arg.AccountInfo.LocateID, account.LocateID)
	require.Equal(t, arg.AccountInfo.Rate, account.Rate)
	require.Equal(t, arg.AccountInfo.Character, account.Character)
	require.Equal(t, arg.AccountInfo.ShowLocate, account.ShowLocate)
	require.Equal(t, arg.AccountInfo.ShowRate, account.ShowRate)
	require.Equal(t, arg.AccountInfo.ShowLocate, account.ShowLocate)
	require.NotEmpty(t, account.CreateAt)
	require.NotEmpty(t, account.UpdateAt)
	require.Equal(t, false, account.IsDelete)

	require.Len(t, tags, len(arg.AccountTechTag))
	require.Len(t, frameworks, len(arg.AccountFrameworkTag))
}

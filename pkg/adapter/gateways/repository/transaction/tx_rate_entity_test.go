package transaction

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/stretchr/testify/require"
)

func TestCreateRateEntityTx(t *testing.T) {
	_, account1 := randomAccount(t)

	arg := repository.CreateRateEntitiesParams{
		AccountID: account1.AccountID,
		Rate:      10,
	}

	require.NoError(t, store.CreateRateEntityTx(context.Background(), arg))

	account2, err := store.GetAccountsByID(context.Background(), account1.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.Rate-account1.Rate, arg.Rate)
}

package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createAccountTest(t *testing.T) Accounts {

	arg := CreateAccountParams{
		UserID:     util.RandomString(8),
		Username:   util.RandomString(8),
		LocateID:   int32(util.Random(47)),
		Rate:       0,
		Email:      util.RandomEmail(),
		ShowRate:   true,
		ShowLocate: true,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.LocateID, account.LocateID)
	require.Equal(t, arg.Rate, account.Rate)
	require.Equal(t, arg.HashedPassword, account.HashedPassword)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.ShowLocate, account.ShowLocate)
	require.Equal(t, arg.ShowRate, account.ShowRate)

	require.NotZero(t, account.CreateAt)
	require.NotZero(t, account.UpdateAt)
	return account
}

func TestCreateAccoutn(t *testing.T) {
	createAccountTest(t)
}

func TestGetAccount(t *testing.T) {
	account := createAccountTest(t)

	locate, err := testQueries.GetLocate(context.Background(), account.LocateID)
	require.NoError(t, err)
	require.NotEmpty(t, locate)

	result, err := testQueries.GetAccount(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, account.UserID, result.UserID)
	require.Equal(t, account.Username, result.Username)
	// locate 変換必要

	require.Equal(t, locate.Name, result.Locate)
	require.Equal(t, account.Rate, result.Rate)
	require.Equal(t, account.ShowLocate, result.ShowLocate)
	require.Equal(t, account.ShowRate, result.ShowRate)

	require.NotZero(t, account.CreateAt)
	require.NotZero(t, account.UpdateAt)
}

func TestListAccount(t *testing.T) {
	n := 10

	for i := 0; i < n; i++ {
		createAccountTest(t)
	}
	arg := ListAccountsParams{
		Limit:  int32(n),
		Offset: 0,
	}

	result, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, n)
}
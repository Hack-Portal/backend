package db

import (
	"context"
	"database/sql"
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

func TestCreateAccount(t *testing.T) {
	createAccountTest(t)
}

func TestGetAccountByID(t *testing.T) {
	account := createAccountTest(t)

	result, err := testQueries.GetAccountByID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	locate, err := testQueries.GetLocateByID(context.Background(), account.LocateID)
	require.NoError(t, err)
	require.NotEmpty(t, locate)

	require.Equal(t, account.UserID, result.UserID)
	require.Equal(t, account.Username, result.Username)
	require.Equal(t, locate.Name, result.Locate)
	require.Equal(t, account.Rate, result.Rate)
	require.Equal(t, account.ShowLocate, result.ShowLocate)
	require.Equal(t, account.ShowRate, result.ShowRate)

	require.NotZero(t, account.CreateAt)
	require.NotZero(t, account.UpdateAt)
}

func TestListAccount(t *testing.T) {
	var lastAccount Accounts
	n := 10

	for i := 0; i < n; i++ {
		lastAccount = createAccountTest(t)
	}

	username := util.Remove5Strings(lastAccount.Username)
	arg := ListAccountsParams{
		Username: "%" + username + "%",
		Limit:    int32(n),
		Offset:   0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}

func TestGetAccountByEmail(t *testing.T) {
	account := createAccountTest(t)

	result, err := testQueries.GetAccountByEmail(context.Background(), account.Email)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	locate, err := testQueries.GetLocateByID(context.Background(), account.LocateID)
	require.NoError(t, err)
	require.NotEmpty(t, locate)

	require.Equal(t, account.UserID, result.UserID)
	require.Equal(t, account.Username, result.Username)
	require.Equal(t, locate.Name, result.Locate)
	require.Equal(t, account.Rate, result.Rate)
	require.Equal(t, account.ShowLocate, result.ShowLocate)
	require.Equal(t, account.ShowRate, result.ShowRate)

	require.NotZero(t, account.CreateAt)
	require.NotZero(t, account.UpdateAt)
}

func TestSoftDeleteAccount(t *testing.T) {
	account1 := createAccountTest(t)

	deletedAccount, err := testQueries.SoftDeleteAccount(context.Background(), account1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAccount)

	account2, err := testQueries.GetAccountByID(context.Background(), account1.UserID)
	require.NoError(t, err)
	require.Empty(t, account2)
}

func TestUpdateAccount(t *testing.T) {
	baseAccount := createAccountTest(t)
	testCase := []struct {
		name      string
		arg       UpdateAccountParams
		checkData func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts)
	}{
		{
			name: "update-username",
			arg: UpdateAccountParams{
				UserID:          baseAccount.UserID,
				Username:        "changed-name",
				Icon:            baseAccount.Icon,
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        baseAccount.LocateID,
				Rate:            baseAccount.Rate,
				HashedPassword:  baseAccount.HashedPassword,
				Email:           baseAccount.Email,
				ShowLocate:      baseAccount.ShowLocate,
				ShowRate:        baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, arg.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-icon",
			arg: UpdateAccountParams{
				UserID:   baseAccount.UserID,
				Username: baseAccount.Username,
				Icon: sql.NullString{
					String: "chenged-icon",
					Valid:  true,
				},
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        baseAccount.LocateID,
				Rate:            baseAccount.Rate,
				HashedPassword:  baseAccount.HashedPassword,
				Email:           baseAccount.Email,
				ShowLocate:      baseAccount.ShowLocate,
				ShowRate:        baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, arg.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-ExplanatoryText",
			arg: UpdateAccountParams{
				UserID:   baseAccount.UserID,
				Username: baseAccount.Username,
				Icon:     baseAccount.Icon,
				ExplanatoryText: sql.NullString{
					String: "changed-explanatoryText",
					Valid:  true,
				},
				LocateID:       baseAccount.LocateID,
				Rate:           baseAccount.Rate,
				HashedPassword: baseAccount.HashedPassword,
				Email:          baseAccount.Email,
				ShowLocate:     baseAccount.ShowLocate,
				ShowRate:       baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, arg.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-LocateID",
			arg: UpdateAccountParams{
				UserID:          baseAccount.UserID,
				Username:        baseAccount.Username,
				Icon:            baseAccount.Icon,
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        1,
				Rate:            baseAccount.Rate,
				HashedPassword:  baseAccount.HashedPassword,
				Email:           baseAccount.Email,
				ShowLocate:      baseAccount.ShowLocate,
				ShowRate:        baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, arg.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-Rate",
			arg: UpdateAccountParams{
				UserID:          baseAccount.UserID,
				Username:        baseAccount.Username,
				Icon:            baseAccount.Icon,
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        baseAccount.LocateID,
				Rate:            10,
				HashedPassword:  baseAccount.HashedPassword,
				Email:           baseAccount.Email,
				ShowLocate:      baseAccount.ShowLocate,
				ShowRate:        baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, arg.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-HashedPassword",
			arg: UpdateAccountParams{
				UserID:          baseAccount.UserID,
				Username:        baseAccount.Username,
				Icon:            baseAccount.Icon,
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        baseAccount.LocateID,
				Rate:            baseAccount.Rate,
				HashedPassword: sql.NullString{
					String: "changed-password",
					Valid:  true,
				},
				Email:      baseAccount.Email,
				ShowLocate: baseAccount.ShowLocate,
				ShowRate:   baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, arg.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-Email",
			arg: UpdateAccountParams{
				UserID:          baseAccount.UserID,
				Username:        baseAccount.Username,
				Icon:            baseAccount.Icon,
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        baseAccount.LocateID,
				Rate:            baseAccount.Rate,
				HashedPassword:  baseAccount.HashedPassword,
				Email:           util.RandomEmail(),
				ShowLocate:      baseAccount.ShowLocate,
				ShowRate:        baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, arg.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-ShowLocate",
			arg: UpdateAccountParams{
				UserID:          baseAccount.UserID,
				Username:        baseAccount.Username,
				Icon:            baseAccount.Icon,
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        baseAccount.LocateID,
				Rate:            baseAccount.Rate,
				HashedPassword:  baseAccount.HashedPassword,
				Email:           baseAccount.Email,
				ShowLocate:      false,
				ShowRate:        baseAccount.ShowRate,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, arg.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, baseAccount.ShowRate, UpdatedAccount.ShowRate)
			},
		},
		{
			name: "update-ShowRate",
			arg: UpdateAccountParams{
				UserID:          baseAccount.UserID,
				Username:        baseAccount.Username,
				Icon:            baseAccount.Icon,
				ExplanatoryText: baseAccount.ExplanatoryText,
				LocateID:        baseAccount.LocateID,
				Rate:            baseAccount.Rate,
				HashedPassword:  baseAccount.HashedPassword,
				Email:           baseAccount.Email,
				ShowLocate:      baseAccount.ShowLocate,
				ShowRate:        false,
			},
			checkData: func(t *testing.T, arg UpdateAccountParams, baseAccount, UpdatedAccount Accounts) {
				require.Equal(t, baseAccount.Username, UpdatedAccount.Username)
				require.Equal(t, baseAccount.Icon, UpdatedAccount.Icon)
				require.Equal(t, baseAccount.ExplanatoryText, UpdatedAccount.ExplanatoryText)
				require.Equal(t, baseAccount.LocateID, UpdatedAccount.LocateID)
				require.Equal(t, baseAccount.Rate, UpdatedAccount.Rate)
				require.Equal(t, baseAccount.HashedPassword, UpdatedAccount.HashedPassword)
				require.Equal(t, baseAccount.Email, UpdatedAccount.Email)
				require.Equal(t, baseAccount.ShowLocate, UpdatedAccount.ShowLocate)
				require.Equal(t, arg.ShowRate, UpdatedAccount.ShowRate)
			},
		},
	}

	for _, tc := range testCase {

		UpdatedAccount, err := testQueries.UpdateAccount(context.Background(), tc.arg)
		require.NoError(t, err)
		require.NotEmpty(t, UpdatedAccount)

		tc.checkData(t, tc.arg, baseAccount, UpdatedAccount)
	}

}

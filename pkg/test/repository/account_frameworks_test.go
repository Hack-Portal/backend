package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func createAccountFrameworksTest(t *testing.T, account repository.Account) repository.AccountFramework {
	frameworks := listFrameworkTest(t)
	randomId := util.Random(len(frameworks) - 1)

	arg := repository.CreateAccountFrameworksParams{
		AccountID:   account.AccountID,
		FrameworkID: frameworks[randomId].FrameworkID,
	}

	accountFrameworks, err := testQueries.CreateAccountFrameworks(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accountFrameworks)

	require.Equal(t, arg.AccountID, accountFrameworks.AccountID)
	require.Equal(t, arg.FrameworkID, accountFrameworks.FrameworkID)

	return accountFrameworks
}

func TestCreateAccountFrameworks(t *testing.T) {
	account := CreateAccountTest(t)
	createAccountFrameworksTest(t, account)
}

func TestListAccountFrameworks(t *testing.T) {
	n := 5
	account := CreateAccountTest(t)
	for i := 0; i < n; i++ {
		createAccountFrameworksTest(t, account)
	}

	listAccountsFramework, err := testQueries.ListAccountFrameworksByUserID(context.Background(), account.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, listAccountsFramework)
	require.Len(t, listAccountsFramework, n)

	for _, accountsFramework := range listAccountsFramework {
		require.NotEmpty(t, accountsFramework)
	}
}

func TestDeleteAccountFrameworksByUserID(t *testing.T) {
	n := 5
	account := CreateAccountTest(t)
	for i := 0; i < n; i++ {
		createAccountFrameworksTest(t, account)
	}

	err := testQueries.DeleteAccountFrameworkByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	accountFramework, err := testQueries.ListAccountFrameworksByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.Len(t, accountFramework, 0)
	for _, af := range accountFramework {
		require.Empty(t, af)
	}
}

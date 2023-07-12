package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createAccountFrameworksTest(t *testing.T, account Accounts) AccountFrameworks {
	frameworks := listFrameworkTest(t)
	randomId := util.Random(len(frameworks) - 1)

	arg := CreateAccountFrameworkParams{
		UserID:      account.UserID,
		FrameworkID: frameworks[randomId].FrameworkID,
	}

	accountFrameworks, err := testQueries.CreateAccountFramework(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accountFrameworks)

	require.Equal(t, arg.UserID, accountFrameworks.UserID)
	require.Equal(t, arg.FrameworkID, accountFrameworks.FrameworkID)

	return accountFrameworks
}

func TestCreateAccountFrameworks(t *testing.T) {
	account := createAccountTest(t)
	createAccountFrameworksTest(t, account)
}

func TestListAccountFrameworks(t *testing.T) {
	n := 5
	account := createAccountTest(t)
	for i := 0; i < n; i++ {
		createAccountFrameworksTest(t, account)
	}

	listAccountsFramework, err := testQueries.ListAccountFrameworksByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, listAccountsFramework)
	require.Len(t, listAccountsFramework, n)

	for _, accountsFramework := range listAccountsFramework {
		require.NotEmpty(t, accountsFramework)
	}
}

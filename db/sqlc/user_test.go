package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

// 新しいランダムなユーザを作る
func createRandomUser(t *testing.T) {
	hashed_password, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUsersParams{
		HashedPassword: hashed_password,
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreateAt)
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

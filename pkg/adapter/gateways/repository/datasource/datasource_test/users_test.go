package repository_test

import (
	"context"
	"database/sql"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func CreateUserTest(t *testing.T) repository.User {

	arg := repository.CreateUsersParams{
		Email:          sql.NullString{String: password.RandomEmail(), Valid: true},
		HashedPassword: sql.NullString{String: password.RandomString(10), Valid: true},
	}

	user, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, false, user.IsDelete)

	require.NotEmpty(t, user.CreateAt)
	require.NotEmpty(t, user.UpdateAt)
	return user
}

func TestCreateUsers(t *testing.T) {
	CreateUserTest(t)
}

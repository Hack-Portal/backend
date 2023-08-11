package interactortest_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	mock_transaction "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/mock"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func TestGetAccountByID(t *testing.T) {
	user := randomUser()
	newAccount := randomAccount(user.UserID)

	testCases := []struct {
		Name          string
		AccountID     string
		buildStubs    func(store *mock_transaction.MockStore)
		checkResponse func(t *testing.T, account domain.AccountResponses)
	}{
		{
			Name:      "success",
			AccountID: newAccount.AccountID,
			buildStubs: func(store *mock_transaction.MockStore) {
				store.EXPECT().
					GetAccountsByID(gomock.Any(), gomock.Eq(newAccount.AccountID)).
					Times(1).
					Return(newAccount, nil)
			},
			checkResponse: func(t *testing.T, account domain.AccountResponses) {
				require.Equal(t, newAccount.AccountID, account.AccountID)
				require.Equal(t, newAccount.Username, account.Username)
				require.Equal(t, newAccount.Icon, account.Icon)
				require.Equal(t, newAccount.ExplanatoryText, account.ExplanatoryText)
				require.Equal(t, newAccount.Rate, account.Rate)
				require.Equal(t, newAccount.ShowLocate, account.ShowLocate)
				require.Equal(t, newAccount.ShowRate, account.ShowRate)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_transaction.NewMockStore(ctrl)
			testCase.buildStubs(store)

			testAccount := usecase.NewAccountUsercase(store, time.Duration(10))

			account, err := testAccount.GetAccountByID(context.Background(), newAccount.AccountID)
			require.NoError(t, err)
			require.NotEmpty(t, account)

			testCase.checkResponse(t, account)
		})
	}

}
func randomUser() repository.User {
	return repository.User{
		UserID:         password.RandomString(10),
		Email:          dbutil.ToSqlNullString(password.RandomEmail()),
		HashedPassword: dbutil.ToSqlNullString(password.RandomString(10)),
		CreateAt:       time.Now().Add(-time.Hour * 24),
		UpdateAt:       time.Now().Add(-time.Hour * 24),
		IsDelete:       false,
	}
}
func randomAccount(userID string) repository.Account {
	return repository.Account{
		AccountID:       password.RandomString(10),
		UserID:          userID,
		Username:        password.RandomString(10),
		Icon:            dbutil.ToSqlNullString(password.RandomString(10)),
		ExplanatoryText: dbutil.ToSqlNullString(password.RandomString(10)),
		LocateID:        int32(password.Random(47)),
		Rate:            int32(password.Random(100)),
		ShowLocate:      true,
		ShowRate:        true,
		CreateAt:        time.Now().Add(-time.Hour * 24),
		UpdateAt:        time.Now().Add(-time.Hour * 24),
		IsDelete:        false,
	}
}

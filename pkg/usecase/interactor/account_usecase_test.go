package usecase

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
	"github.com/stretchr/testify/require"
)

func newAccountUsecase() inputport.AccountUsecase {
	return NewAccountUsercase(store, time.Duration(time.Second*5))
}

func testImage(t *testing.T) []byte {
	image, err := ioutil.ReadFile("../../../color.png")
	require.NoError(t, err)
	require.NotEmpty(t, image)
	return image
}

func createAccountTest(t *testing.T, au inputport.AccountUsecase) (domain.AccountResponses, string) {

	image := testImage(t)
	email := util.RandomEmail()
	arg := domain.CreateAccount{
		ReqBody: domain.CreateAccountRequest{
			AccountID:       util.RandomString(10),
			Username:        util.RandomString(10),
			ExplanatoryText: util.RandomString(10),
			LocateID:        int32(util.Random(46) + 1),
			ShowLocate:      false,
			ShowRate:        false,
		},
		TechTags:   util.RandomSelection(13, 3),
		Frameworks: util.RandomSelection(50, 3),
	}

	account, err := au.CreateAccount(context.Background(), arg, image, email)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	locate, err := store.GetLocatesByID(context.Background(), arg.ReqBody.LocateID)
	require.NoError(t, err)
	require.NotEmpty(t, locate)

	require.Equal(t, arg.ReqBody.AccountID, account.AccountID)
	require.Equal(t, arg.ReqBody.Username, account.Username)
	require.Equal(t, arg.ReqBody.ExplanatoryText, account.ExplanatoryText)
	require.Equal(t, arg.ReqBody.ShowLocate, account.ShowLocate)
	require.Equal(t, arg.ReqBody.ShowRate, account.ShowRate)

	require.Len(t, account.TechTags, len(arg.TechTags))
	require.Len(t, account.Frameworks, len(arg.Frameworks))
	return account, email
}
func TestCreateAccout(t *testing.T) {
	au := newAccountUsecase()
	createAccountTest(t, au)
}

func TestGetAccountByID(t *testing.T) {
	au := newAccountUsecase()
	account, _ := createAccountTest(t, au)

	getAccount, err := au.GetAccountByID(context.Background(), account.AccountID, nil)
	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, account, getAccount)
}

func TestGetAccountByEmail(t *testing.T) {
	au := newAccountUsecase()
	account, email := createAccountTest(t, au)

	getAccount, err := au.GetAccountByEmail(context.Background(), email)
	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, account, getAccount)
}

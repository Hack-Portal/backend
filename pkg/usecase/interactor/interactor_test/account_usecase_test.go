package interactortest_test

import (
	"testing"

	mock_transaction "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/mock"
)

func TestGetAccountByID(t *testing.T) {

	testCase := []struct {
		Name          string
		AccountID     string
		buildStubs    func(store *mock_transaction.MockStore)
		checkResponse func(t *testing.T)
	}{
		{
			Name: "Success",
		},
	}
}

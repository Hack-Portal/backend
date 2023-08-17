package transaction

import (
	"testing"

	mock "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/mock"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

func TestCreateAccountTx(t *testing.T) {
	testCases := []struct {
		Name          string
		args          domain.CreateAccountParams
		buildStubs    func(store *mock.MockStore)
		checkResponse func(t *testing.T)
	}{}
	print(testCases)
}

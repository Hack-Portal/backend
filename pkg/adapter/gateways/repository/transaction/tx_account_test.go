package transaction

import (
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

func TestCreateAccountTx(t *testing.T) {
	testCases := []struct {
		name          string
		arg           domain.CreateAccountParams
		checkReturens func(t *testing.T)
	}{
		{
			name: "success",
			arg: domain.CreateAccountParams{
				AccountInfo:         repository.CreateAccountsParams{},
				AccountTechTag:      []int32{},
				AccountFrameworkTag: []int32{},
			},
			checkReturens: func(t *testing.T) {},
		},
	}

}

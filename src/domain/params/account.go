package params

import (
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type CreateAccount struct {
	AccountInfo         repository.CreateAccountsParams
	AccountTechTag      []int32
	AccountFrameworkTag []int32
}

type UpdateAccount struct {
	AccountInfo         repository.Account
	AccountTechTag      []int32
	AccountFrameworkTag []int32
}

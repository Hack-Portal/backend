package params

import (
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
)

type CreateAccountParams struct {
	AccountInfo         repository.CreateAccountsParams
	AccountTechTag      []int32
	AccountFrameworkTag []int32
}

type UpdateAccountParams struct {
	AccountInfo         repository.Account
	AccountTechTag      []int32
	AccountFrameworkTag []int32
}
type CreateAccount struct {
	ReqBody    request.CreateAccountRequest
	TechTags   []int32
	Frameworks []int32
}

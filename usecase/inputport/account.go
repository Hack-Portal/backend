package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/domain"
)

type AccountUsecase interface {
	GetAccountByID(ctx context.Context, id string) (domain.AccountResponses, error)
	GetAccountByEmail(ctx context.Context, email string) (domain.AccountResponses, error)
	CreateAccount(ctx context.Context, body domain.CreateAccountTxParams) (domain.AccountResponses, error)
	UpdateAccount(ctx context.Context, body domain.UpdateAccountTxParams) (domain.AccountResponses, error)
	DeleteAccount(ctx context.Context, id string) error
	UploadImage(ctx context.Context, body []byte) (string, error)
}

package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
)

type AccountUsecase interface {
	GetAccountByID(ctx context.Context, id string, token *jwt.FireBaseCustomToken) (domain.AccountResponses, error)
	GetAccountByEmail(ctx context.Context, email string) (domain.AccountResponses, error)
	CreateAccount(ctx context.Context, body domain.CreateAccount, image []byte, email string) (domain.AccountResponses, error)
	UpdateAccount(ctx context.Context, body domain.UpdateAccountParam, image []byte) (domain.AccountResponses, error)
	DeleteAccount(ctx context.Context, id string) error
}

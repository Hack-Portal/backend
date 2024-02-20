package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/jwt"
	"github.com/hackhack-Geek-vol6/backend/src/domain/params"
	"github.com/hackhack-Geek-vol6/backend/src/domain/response"
)

type AccountUsecase interface {
	GetAccountByID(ctx context.Context, id string, token *jwt.FireBaseCustomToken) (response.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (response.Account, error)
	CreateAccount(ctx context.Context, body params.CreateAccount, image []byte) (response.Account, error)
	UpdateAccount(ctx context.Context, body params.UpdateAccount, image []byte) (response.Account, error)
	DeleteAccount(ctx context.Context, id string) error
	GetJoinRoom(ctx context.Context, accountID string) (result []response.GetJoinRoom, err error)
}

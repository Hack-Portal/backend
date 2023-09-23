package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
)

type AccountUsecase interface {
	GetAccountByID(ctx context.Context, id string, token *jwt.FireBaseCustomToken) (response.AccountResponse, error)
	GetAccountByEmail(ctx context.Context, email string) (response.AccountResponse, error)
	CreateAccount(ctx context.Context, body params.CreateAccount, image []byte) (response.AccountResponse, error)
	UpdateAccount(ctx context.Context, body params.UpdateAccount, image []byte) (response.AccountResponse, error)
	DeleteAccount(ctx context.Context, id string) error
	GetJoinRoom(ctx context.Context, accountID string) (result []response.GetJoinRoomResponse, err error)
}

package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type userUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewUserUsercase(store transaction.Store, timeout time.Duration) inputport.UserUsecase {
	return &userUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) CreateUser(ctx context.Context, body repository.CreateUsersParams) (user repository.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, uu.contextTimeout)
	defer cancel()

	user, err = uu.store.CreateUsers(ctx, body)
	return
}

package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	tokens "github.com/hackhack-Geek-vol6/backend/pkg/util/token"
)

type userUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
	tokenMaker     tokens.Maker
}

func NewUserUsercase(store transaction.Store, timeout time.Duration, tokenMaker tokens.Maker) inputport.UserUsecase {
	return &userUsecase{
		store:          store,
		contextTimeout: timeout,
		tokenMaker:     tokenMaker,
	}
}

func (uu *userUsecase) CreateUser(ctx context.Context, arg domain.CreateUserRequest, duration time.Duration) (result domain.CreateUserResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, uu.contextTimeout)
	defer cancel()

	hashedPassword, err := password.HashPassword(arg.Password)
	if err != nil {
		return
	}

	user, err := uu.store.CreateUsers(ctx, repository.CreateUsersParams{
		UserID:         uuid.New().String(),
		Email:          dbutil.ToSqlNullString(arg.Email),
		HashedPassword: dbutil.ToSqlNullString(hashedPassword),
	})

	token, err := uu.tokenMaker.CreateToken(arg.Email, duration)
	if err != nil {
		return
	}

	result = domain.CreateUserResponse{
		UserID:    user.UserID,
		Token:     token,
		CreatedAt: user.CreateAt,
	}

	return
}

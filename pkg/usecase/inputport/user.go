package inputport

import (
	"context"
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, arg domain.CreateUserRequest, duration time.Duration) (result domain.CreateUserResponse, err error)
}

package auth

import (
	"log"

	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/Hack-Portal/backend/src/utils/password"
	"github.com/labstack/echo/v4"
)

const (
	AuthGuestID int = 3
)

type basicAuth struct {
	userRepo dai.UsersDai
}

func NewBasicAuth(repo dai.UsersDai) Auth {
	return &basicAuth{
		userRepo: repo,
	}
}

func (ba *basicAuth) AuthN() echo.MiddlewareFunc {
	return ba.basic
}

func (ba *basicAuth) basic(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid, pass, ok := c.Request().BasicAuth()
		if !ok {
			// TODO:ここでログを出力する
			// 未認証ユーザ(guest)として扱う
			c.Set(RequestRoleID, AuthGuestID)
			return next(c)
		}
		user, err := ba.userRepo.FindByID(c.Request().Context(), uid)
		if err != nil {
			// TODO:ここでログを出力する
			return echo.ErrInternalServerError
		}
		if user == nil || !user.DeletedAt.IsZero() {
			return echo.ErrUnauthorized
		}

		if err := password.CheckPassword(pass, user.Password); err != nil {
			log.Println(err, user.Password, pass)
			return echo.ErrUnauthorized
		}
		c.Set(RequestUserID, user.UserID)
		c.Set(RequestRoleID, user.Role)

		return next(c)
	}
}

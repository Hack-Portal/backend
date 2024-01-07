package auth

import "github.com/labstack/echo/v4"

type Auth interface {
	AuthN() echo.MiddlewareFunc
}

const (
	IsGuest       = "is_guest"
	RequestUserID = "request_user_id"
	RequestRoleID = "request_role_id"
)

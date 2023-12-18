package middleware

import (
	"github.com/labstack/echo/v4"
)

type EchoMiddleware interface {
	Auth() echo.MiddlewareFunc
	AccessLog() echo.MiddlewareFunc
}

func NewEchoMiddleware() EchoMiddleware {
	return &middleware{}
}

func (m *middleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			panic("implement me")
		}
	}
}

func (m *middleware) AccessLog() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			panic("implement me")
		}
	}
}

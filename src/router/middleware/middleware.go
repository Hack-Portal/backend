package middleware

import (
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/router/middleware/auth"
	"gorm.io/gorm"
)

type middleware struct {
	auth.Auth
}

func NewMiddleware(db *gorm.DB) *middleware {
	return &middleware{
		Auth: auth.NewBasicAuth(gateways.NewUserGateway(db)),
	}
}

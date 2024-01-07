package middleware

import (
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/router/middleware/auth"
	"github.com/Hack-Portal/backend/src/router/middleware/casbin"
	"gorm.io/gorm"
)

type middleware struct {
	auth.Auth
	casbin.RBACPolicy
}

func NewMiddleware(db *gorm.DB) *middleware {
	return &middleware{
		Auth:       auth.NewBasicAuth(gateways.NewUserGateway(db)),
		RBACPolicy: casbin.NewRBAC(gateways.NewRbacPolicyGateway(db)),
	}
}

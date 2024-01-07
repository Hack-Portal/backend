package middleware

import "github.com/Hack-Portal/backend/src/router/middleware/auth"

type Middleware interface {
	auth.Auth
}

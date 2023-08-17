package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
	tokens "github.com/hackhack-Geek-vol6/backend/pkg/util/token"
)

const (
	AuthorizationHeaderKey = "dbauthorization"
	AuthorizationType      = "dbauthorization_type"
	AuthorizationClaimsKey = "authorization_claim"

	GoogleLogin = "google"
	EmailLogin  = "email"
)

func AuthMiddleware(tokenMaker tokens.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtType := ctx.GetHeader(AuthorizationType)
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)

		if len(jwtType) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 1 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		switch jwtType {

		case GoogleLogin: //Emailでのログインtoken
			payload, err := googleLogin(fields)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			ctx.Set(AuthorizationClaimsKey, payload)

		case EmailLogin: //Emailでのログインtoken
			payload, err := emailLogin(tokenMaker, fields)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			ctx.Set(AuthorizationClaimsKey, payload)
		}
		ctx.Next()
	}
}

func googleLogin(fields []string) (*tokens.Payload, error) {
	accessToken := fields[0]
	hCS, err := jwt.JwtDecode.DecomposeFB(accessToken)
	if err != nil {
		return nil, err
	}

	payload, err := jwt.JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		return nil, err
	}
	return &tokens.Payload{
		UserID: payload.UID,
		Email:  payload.Email,
	}, nil
}

func emailLogin(tokenMaker tokens.Maker, fields []string) (*tokens.Payload, error) {
	accessToken := fields[0]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

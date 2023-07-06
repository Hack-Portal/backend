package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const (
	authorizationHeader = "Authorization"
	authorizationType   = "bearer"
	valName             = "token"
)

type FirebaseAuthMiddleware struct {
	client       *auth.Client
	unAuthorized func(ctx *gin.Context)
}

func New(credFileName string, unAuthorized func(ctx *gin.Context)) (*FirebaseAuthMiddleware, error) {
	opt := option.WithCredentialsFile(credFileName)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return &FirebaseAuthMiddleware{
		client:       auth,
		unAuthorized: unAuthorized,
	}, nil
}
func (firebaseAuthMiddleware *FirebaseAuthMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get(authorizationHeader)
		token := strings.Replace(authHeader, authorizationType, "", 1)
		idToken, err := firebaseAuthMiddleware.client.VerifyIDToken(context.Background(), token)
		if err != nil {
			if firebaseAuthMiddleware.unAuthorized != nil {
				firebaseAuthMiddleware.unAuthorized(ctx)
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		ctx.Set(valName, idToken)
		ctx.Next()
	}
}

func ExtractClaims(c *gin.Context) *auth.Token {
	idToken, ok := c.Get(valName)
	if !ok {
		return new(auth.Token)
	}
	return idToken.(*auth.Token)
}

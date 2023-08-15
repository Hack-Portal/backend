package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/middleware"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	_ "github.com/hackhack-Geek-vol6/backend/pkg/docs"
)

func setupCors(router *gin.Engine) {
	router.Use(cors.Default())
}

func Setup(env *bootstrap.Env, timeout time.Duration, store transaction.Store, gin *gin.Engine) {
	setupCors(gin)

	publicRouter := gin.Group("/v1")
	// All Public APIs
	NewEtcRouter(env, timeout, store, publicRouter)
	NewHackathonRouter(env, timeout, store, publicRouter)

	protectRouter := gin.Group("/v1").Use(middleware.AuthMiddleware())
	//TODO:middlewareの追加
	// All Protect APIs
	NewAccountRouter(env, timeout, store, protectRouter)
	NewLikeRouter(env, timeout, store, protectRouter)
	NewFollowRouter(env, timeout, store, protectRouter)
	NewRateRouter(env, timeout, store, protectRouter)
	NewRoomRouter(env, timeout, store, protectRouter)
}

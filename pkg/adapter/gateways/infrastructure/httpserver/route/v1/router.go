package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/hackhack-Geek-vol6/backend/docs"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/apm"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/middleware"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"go.uber.org/zap"
)

func setupCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "dbauthorization")
	router.Use(cors.New(config))
}

func Setup(env *bootstrap.Env, timeout time.Duration, store transaction.Store, gin *gin.Engine, logger *zap.Logger) {
	setupCors(gin)
	gin.Use(middleware.Logger(logger), nrgin.Middleware(apm.NewApm(env)))

	publicRouter := gin.Group("/v1").Use(middleware.CheckJWT())
	// All Public APIs
	NewEtcRouter(env, timeout, store, publicRouter)
	NewHackathonRouter(env, timeout, store, publicRouter)
	setupSwagger(publicRouter)

	protectRouter := gin.Group("/v1").Use(middleware.AuthMiddleware())
	// All Protect APIs
	NewAccountRouter(env, timeout, store, protectRouter, publicRouter)
	NewLikeRouter(env, timeout, store, protectRouter)
	NewPastWorkRouter(env, timeout, store, protectRouter, publicRouter)
	NewFollowRouter(env, timeout, store, protectRouter)
	NewRateRouter(env, timeout, store, protectRouter, publicRouter)
	NewRoomRouter(env, timeout, store, protectRouter)
}

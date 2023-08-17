package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/apm"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/middleware"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	_ "github.com/hackhack-Geek-vol6/backend/pkg/docs"
	tokens "github.com/hackhack-Geek-vol6/backend/pkg/util/token"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

func setupCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{
		"dbauthorization",
		"dbauthorization_type",
	}
	router.Use(cors.New(config))
}

func Setup(env *bootstrap.Env, tokenMaker tokens.Maker, timeout time.Duration, store transaction.Store, gin *gin.Engine) {
	gin.Use(nrgin.Middleware(apm.ApmSetup(env)))
	setupCors(gin)

	publicRouter := gin.Group("/v1")
	// All Public APIs
	NewEtcRouter(env, timeout, store, publicRouter)
	NewHackathonRouter(env, timeout, store, publicRouter)
	NewUserRouter(env, tokenMaker, timeout, store, publicRouter)

	protectRouter := gin.Group("/v1").Use(middleware.AuthMiddleware(tokenMaker))
	//TODO:middlewareの追加
	// All Protect APIs
	NewAccountRouter(env, timeout, store, protectRouter)
	NewLikeRouter(env, timeout, store, protectRouter)
	NewFollowRouter(env, timeout, store, protectRouter)
	NewRateRouter(env, timeout, store, protectRouter)
	NewRoomRouter(env, timeout, store, protectRouter)
}

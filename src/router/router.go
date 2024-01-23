package router

import (
	"github.com/Hack-Portal/backend/src/adapters/controllers"
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/router/middleware"
	v1 "github.com/Hack-Portal/backend/src/router/v1"
	"github.com/Hack-Portal/backend/src/usecases/interactors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type router struct {
	engine *echo.Echo
	db     *gorm.DB
	cache  *redis.Client
	client *s3.Client
	nrApp  *newrelic.Application

	config Config
}

func NewRouter(config Config, db *gorm.DB, cache *redis.Client, nrApp *newrelic.Application, client *s3.Client) *echo.Echo {
	router := &router{
		engine: echo.New(),
		db:     db,
		cache:  cache,
		client: client,
		nrApp:  nrApp,

		config: config,
	}

	router.setup()

	return router.engine
}

func (r *router) setup() {
	// di
	uc := controllers.NewUserController(
		interactors.NewUserInteractor(
			gateways.NewUserGateway(r.db),
			gateways.NewRoleGateway(r.db),
			presenters.NewUserPresenter(),
		),
	)

	// setup newrelic middleware
	r.engine.Use(nrecho.Middleware(r.nrApp))
	// health check
	r.engine.GET("/health", func(c echo.Context) error { return c.String(200, "ok") })
	r.engine.GET("/version", func(c echo.Context) error { return c.String(200, r.config.Version) })
	r.engine.POST("/init_admin", uc.InitAdmin)
	r.engine.POST("/login", uc.Login)

	// status tag
	middleware := middleware.NewMiddleware(r.db)
	{
		v1group := r.engine.Group("/v1")
		v1group.Use(
			middleware.AuthN(),
			middleware.RBACPermission(),
		)

		v1.NewV1Router(v1group, r.db, r.cache, r.client)
	}
}

package router

import (
	"github.com/Hack-Portal/backend/src/router/middleware"
	v1 "github.com/Hack-Portal/backend/src/router/v1"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type router struct {
	engine *echo.Echo
	db     *gorm.DB
	client *s3.Client

	config Config
}

func NewRouter(config Config, db *gorm.DB, client *s3.Client) *echo.Echo {
	router := &router{
		engine: echo.New(),
		db:     db,
		client: client,
		config: config,
	}

	router.setup()

	return router.engine
}

func (r *router) setup() {
	// di

	// health check
	r.engine.GET("/health", func(c echo.Context) error { return c.String(200, "ok") })
	r.engine.GET("/version", func(c echo.Context) error { return c.String(200, r.config.Version) })

	// status tag
	middleware := middleware.NewMiddleware(r.db)
	{
		v1group := r.engine.Group("/v1")
		v1group.Use(
			middleware.AuthN(),
		)

		v1.NewV1Router(v1group)
	}
}

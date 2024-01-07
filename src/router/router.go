package router

import (
	"github.com/Hack-Portal/backend/src/router/middleware/casbin"
	v1 "github.com/Hack-Portal/backend/src/router/v1"
	"github.com/labstack/echo/v4"
)

type router struct {
	engine *echo.Echo

	config Config
}

func NewRouter(config Config) *echo.Echo {
	router := &router{
		engine: echo.New(),
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
	// authorization
	{
		v1group := r.engine.Group("/v1")
		v1group.Use(casbin.Authorization())
		v1.NewV1Router(v1group)
	}
}

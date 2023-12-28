package echo

import (
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type echoServer struct {
	engine *echo.Echo
	v1     *echo.Group

	app    *newrelic.Application
	db     *gorm.DB
	client *s3.Client
	logger *slog.Logger
}

func NewEchoServer(db *gorm.DB, client *s3.Client, app *newrelic.Application, logger *slog.Logger) *echo.Echo {
	router := &echoServer{
		engine: echo.New(),
		app:    app,
		client: client,
		db:     db,
	}

	router.setupMiddleware()

	router.v1 = router.engine.Group("/v1")
	router.StatusTag()
	router.Hackathon()
	// TODO: setup routing
	// router.Proposal()
	// router.Hackathon()

	router.engine.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	return router.engine
}

func (es *echoServer) setupMiddleware() {
	// New Relic Setting
	es.engine.Use(nrecho.Middleware(es.app))
	// Cross-Origin Resource Sharing
	es.engine.Use(
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}),
	)
}

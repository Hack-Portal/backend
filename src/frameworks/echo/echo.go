package echo

import (
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	engine *echo.Echo
	v1     *echo.Group

	db     *gorm.DB
	client *s3.Client
	logger *slog.Logger
}

func NewEchoServer(db *gorm.DB, client *s3.Client, logger *slog.Logger) *echo.Echo {
	router := &echoServer{
		engine: echo.New(),
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
	es.engine.Use(
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}),
	)
	es.engine.Use(
		echoMiddleware.RequestLoggerWithConfig(
			echoMiddleware.RequestLoggerConfig{
				LogStatus:    true,
				LogURI:       true,
				LogLatency:   true,
				LogProtocol:  true,
				LogRemoteIP:  true,
				LogMethod:    true,
				LogURIPath:   true,
				LogRoutePath: true,
				LogUserAgent: true,
				LogError:     true,
				LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
					return nil
				},
			},
		),
	)
}

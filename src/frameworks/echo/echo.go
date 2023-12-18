package echo

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	engine *echo.Echo
	v1     *echo.Group

	db     *gorm.DB
	logger *slog.Logger
}

func NewEchoServer(db *gorm.DB, logger *slog.Logger) *echo.Echo {
	router := &echoServer{
		engine: echo.New(),
		db:     db,
	}

	router.setupMiddleware()

	router.v1 = router.engine.Group("/v1")
	router.StatusTag()
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
}

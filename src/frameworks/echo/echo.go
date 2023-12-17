package echo

import (
	"log/slog"

	"github.com/labstack/echo/v4"
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
	}

	router.setupMiddleware()

	router.v1 = router.engine.Group("/v1")
	// TODO: setup routing
	// router.Proposal()
	// router.Hackathon()
	// router.StatusTag()

	return router.engine
}

func (es *echoServer) setupMiddleware() {

}

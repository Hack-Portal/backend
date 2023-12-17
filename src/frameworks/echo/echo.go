package echo

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type echoServer struct {
	engine *echo.Echo

	db     *gorm.DB
	logger *slog.Logger
}

func NewEchoServer(db *gorm.DB, logger *slog.Logger) *echo.Echo {
	router := &echoServer{
		engine: echo.New(),
	}

	router.setupMiddleware()

	// TODO: setup routing
	// router.Proposal()
	// router.Hackathon()
	// router.StatusTag()

	return router.engine
}

func (es *echoServer) setupMiddleware() {

}

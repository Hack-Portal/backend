package echo

import "github.com/labstack/echo/v4"

type echoServer struct {
	engine *echo.Echo
}

func NewEchoServer() *echo.Echo {
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

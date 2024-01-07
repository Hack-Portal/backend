package v1

import "github.com/labstack/echo/v4"

type v1router struct {
	v1 *echo.Group
}

func NewV1Router(e *echo.Group) {
	router := &v1router{
		v1: e,
	}

	router.statusTag()
	router.hackathon()
	return
}

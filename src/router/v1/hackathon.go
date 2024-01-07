package v1

import "github.com/labstack/echo/v4"

func (r *v1router) hackathon() {
	hackathon := r.v1.Group("/hackathons")

	hackathon.POST("", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	hackathon.GET("", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	hackathon.PUT("/:hackathon_id", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	hackathon.DELETE("/:hackathon_id", func(c echo.Context) error {
		return c.String(200, "ok")
	})
}

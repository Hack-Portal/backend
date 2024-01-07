package v1

import "github.com/labstack/echo/v4"

func (r *v1router) statusTag() {
	statusTag := r.v1.Group("/status_tags")

	statusTag.GET("", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	statusTag.POST("", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	statusTag.PUT("/:hackathon_id", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	statusTag.DELETE("/:hackathon_id", func(c echo.Context) error {
		return c.String(200, "ok")
	})
}

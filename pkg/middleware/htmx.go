// ==================================================================
// Custom middleware for HTMX requests
// ==================================================================

package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTMXGuard is a middleware that blocks non-HTMX requests
// Use this to block direct browser access to an endpoint
func HTMXGuard() echo.MiddlewareFunc {
	return blockNonHTMX
}

func blockNonHTMX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hxRequest := c.Request().Header.Get("hx-request")
		if hxRequest == "" {
			return c.HTML(http.StatusGone, "This endpoint only accepts HX requests")
		}

		return next(c)
	}
}

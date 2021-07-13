package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) handleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

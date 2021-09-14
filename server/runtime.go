package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func runtime() (e *echo.Echo) {
	e = echo.New()
	e.Any("/", health)
	return
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func App(c echo.Context) error {
	return c.Render(http.StatusOK, "app", nil)
}

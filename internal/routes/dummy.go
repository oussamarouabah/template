package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func dummy(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "hello world!"})
}

package userController

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "id")
}

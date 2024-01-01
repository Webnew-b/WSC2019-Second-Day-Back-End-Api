package attendeesController

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/tools/logUtil"
)

func handleAttendeesError(err error) error {
	logUtil.Log.Println(err.Error())
	switch {

	case errors.Is(err, &attendeesError.NotFound{}):
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

	case errors.Is(err, &attendeesError.LoginKeyNotExist{}):
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

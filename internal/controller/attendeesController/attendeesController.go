package attendeesController

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/tools/logUtil"
)

var valid = validator.New()

func validate(i interface{}, errMsg string) error {
	err := valid.Struct(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errMsg)
	}
	return nil
}

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

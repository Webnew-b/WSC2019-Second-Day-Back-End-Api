package attendeesController

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/service/attendeesService"
	"wscmakebygo.com/tools"
)

var valid = validator.New()

func validate(i interface{}, errMsg string) error {
	err := valid.Struct(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errMsg)
	}
	return nil
}

func AttendeesLogin(c echo.Context) error {
	user := new(api.LoginRequest)

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, attendeesError.ErrInvalidLoginMessage)
	}

	if err := validate(user, attendeesError.ErrInvalidLoginMessage); err != nil {
		return err
	}

	attendees, err := attendeesService.AttendeesLogin(*user)

	if err != nil {
		return handleAttendeesError(err)
	}
	return c.JSON(http.StatusOK, attendees)
}

func AttendeesLogout(c echo.Context) error {
	req := api.LogoutRequest{
		Token: c.QueryParam("token"),
	}

	if err := validate(req, attendeesError.ErrInvalidTokenMessage); err != nil {
		return err
	}

	msg, err := attendeesService.AttendeesLogout(req)
	if err != nil {
		return handleAttendeesError(err)
	}

	return c.JSON(http.StatusOK, &api.LogoutRes{
		Message: msg,
	})
}

func handleAttendeesError(err error) error {
	tools.Log.Println(err.Error())
	switch {

	case errors.Is(err, &attendeesError.NotFound{}):
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

	case errors.Is(err, &attendeesError.LoginKeyNotExist{}):
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

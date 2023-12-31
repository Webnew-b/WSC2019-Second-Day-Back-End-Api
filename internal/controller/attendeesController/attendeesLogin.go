package attendeesController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/service/attendeesService"
)

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

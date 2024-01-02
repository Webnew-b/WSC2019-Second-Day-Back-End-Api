package attendeesController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/service/attendeesService"
)

func AttendeesLogout(c echo.Context) error {
	req := api.LogoutRequest{
		Token: c.QueryParam("token"),
	}

	if err := apperrors.ValidateStruct(req, attendeesError.ErrInvalidTokenMessage); err != nil {
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

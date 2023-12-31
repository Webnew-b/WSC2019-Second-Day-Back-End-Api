package eventsController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/service/eventService"
)

func EventReg(c echo.Context) error {
	param := api.EventRegRequestParams{
		OrgSlug: c.Param("organizerSlug"),
		EvSlug:  c.Param("eventSlug"),
		Token:   c.QueryParam("token"),
	}
	body := new(api.EventRegRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, eventError.INVALID_BODY)
	}

	if err := validate(param, eventError.INVALID_PARAMS); err != nil {
		return err
	}
	if err := validate(body, eventError.INVALID_BODY); err != nil {
		return err
	}

	params := api.EventRegParams{
		EventRegRequestParams: &param,
		EventRegRequestBody:   body,
	}

	res, err := eventService.RegEvent(&params)
	if err != nil {
		return handleEventError(err)
	}
	return c.JSON(http.StatusOK, res)
}

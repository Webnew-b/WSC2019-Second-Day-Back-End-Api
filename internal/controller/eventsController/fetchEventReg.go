package eventsController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/service/eventService"
)

func FetchEventReg(c echo.Context) error {
	fetchEventRegReq := api.FetchEventRegReq{
		Token: c.QueryParam("token"),
	}
	err := validate(fetchEventRegReq, eventError.INVALID_PARAMS)
	if err != nil {
		return handleEventError(err)
	}

	res, err := eventService.FetchEventRegDetail(&fetchEventRegReq)
	if err != nil {
		return handleEventError(err)
	}
	return c.JSON(http.StatusOK, res)
}

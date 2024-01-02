package eventsController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/internal/service/eventService"
)

func FetchEventReg(c echo.Context) error {
	var (
		attendee model.Attendees
		res      *api.FetchEventRegRes
		err      error
	)

	value := c.Get("attendee")
	attendee, err = attendee.CheckAttendeeType(value)
	if err != nil {
		return err
	}

	fetchEventRegReq := api.FetchEventRegReq{
		Id: attendee.ID,
	}

	err = apperrors.ValidateStruct(fetchEventRegReq, eventError.INVALID_PARAMS)
	if err != nil {
		return handleEventError(err)
	}

	res, err = eventService.FetchEventRegDetail(&fetchEventRegReq)
	if err != nil {
		return handleEventError(err)
	}
	return c.JSON(http.StatusOK, res)
}

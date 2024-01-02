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

func EventReg(c echo.Context) error {
	var (
		attendee model.Attendees
		err      error
	)

	param := api.EventRegRequestParams{
		OrgSlug: c.Param("organizerSlug"),
		EvSlug:  c.Param("eventSlug"),
	}
	body := new(api.EventRegRequestBody)

	value := c.Get("attendee")
	attendee, err = attendee.CheckAttendeeType(value)
	if err != nil {
		return err
	}

	if err = c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, eventError.INVALID_BODY)
	}

	if err = apperrors.ValidateStruct(param, eventError.INVALID_PARAMS); err != nil {
		return err
	}
	if err = apperrors.ValidateStruct(body, eventError.INVALID_BODY); err != nil {
		return err
	}

	params := api.EventRegParams{
		AttendeeId:            attendee.ID,
		EventRegRequestParams: &param,
		EventRegRequestBody:   body,
	}

	res, err := eventService.RegEvent(&params)
	if err != nil {
		return handleEventError(err)
	}
	return c.JSON(http.StatusOK, res)
}

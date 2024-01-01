package eventsController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/internal/apperrors"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/params/eventParams"
	"wscmakebygo.com/internal/service/eventService"
)

func GetEventDetail(c echo.Context) error {
	eventFetchRequest := eventParams.EventFetchRequest{
		OrgSlug: c.Param("organizerSlug"),
		EvSlug:  c.Param("eventSlug"),
	}
	err := apperrors.ValidateStruct(eventFetchRequest, eventError.EVENT_NOT_FOUND)
	if err != nil {
		return handleEventError(err)
	}

	res, err := eventService.FetchEventDetail(eventFetchRequest)
	if err != nil {
		return handleEventError(err)
	}

	return c.JSON(http.StatusOK, res)
}

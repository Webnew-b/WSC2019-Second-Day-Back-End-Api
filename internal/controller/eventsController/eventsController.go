package eventsController

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/apperrors/organizerError"
	"wscmakebygo.com/internal/controller"
	"wscmakebygo.com/internal/params/eventParams"
	"wscmakebygo.com/internal/service/eventService"
	"wscmakebygo.com/tools"
)

func GetEvents(c echo.Context) error {
	events, err := eventService.GetAllEventAndOrganizer()
	if err != nil {
		return err
	}
	res := api.EventsRes{
		Events: *events,
	}
	return c.JSON(http.StatusOK, res)
}

func GetEventDetail(c echo.Context) error {
	eventFetchRequest := eventParams.EventFetchRequest{
		OrgSlug: c.Param("organizerSlug"),
		EvSlug:  c.Param("eventSlug"),
	}
	err := controller.GetValidator().Struct(eventFetchRequest)
	if err != nil {
		return handleEventDetailError(err)
	}

	res, err := eventService.FetchEventDetail(eventFetchRequest)
	if err != nil {
		return handleEventDetailError(err)
	}

	return c.JSON(http.StatusOK, res)
}

func handleEventDetailError(err error) error {
	tools.Log.Println(err.Error())
	switch {
	case errors.Is(err, &eventError.EventSlugNotFoundError{}),
		errors.Is(err, &organizerError.OrganizerSlugNotFoundError{}):
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

package eventsController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/service/eventService"
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

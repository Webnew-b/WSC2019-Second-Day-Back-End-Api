package eventsController

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
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
	err := controller.Validator.Struct(eventFetchRequest)
	if err != nil {
		//todo 需要做一个全体的错误处理
		tools.Log.Println(err.Error())
		return err
	}
	res, err := eventService.FetchEventDetail(eventFetchRequest)
	if err != nil {
		tools.Log.Println(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, res)
}

package eventsController

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/apperrors/organizerError"
	"wscmakebygo.com/internal/apperrors/registrationsError"
	"wscmakebygo.com/internal/apperrors/ticketsError"
	"wscmakebygo.com/internal/params/eventParams"
	"wscmakebygo.com/internal/service/eventService"
	"wscmakebygo.com/tools"
)

var valid = validator.New()

func validate(i interface{}, errMsg string) error {
	err := valid.Struct(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errMsg)
	}
	return nil
}

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
	err := validate(eventFetchRequest, eventError.EVENT_NOT_FOUND)
	if err != nil {
		return handleEventError(err)
	}

	res, err := eventService.FetchEventDetail(eventFetchRequest)
	if err != nil {
		return handleEventError(err)
	}

	return c.JSON(http.StatusOK, res)
}

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

func handleEventError(err error) error {
	tools.Log.Println(err.Error())
	switch {
	case errors.Is(err, &eventError.EventSlugNotFoundError{}),
		errors.Is(err, &organizerError.OrganizerSlugNotFoundError{}),
		errors.Is(err, &attendeesError.NotLogin{}),
		errors.Is(err, &ticketsError.NotAvailable{}),
		errors.Is(err, &registrationsError.AlreadyRegistrar{}):
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

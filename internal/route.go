package internal

import (
	"github.com/labstack/echo/v4"
	"wscmakebygo.com/global/route"
	"wscmakebygo.com/internal/controller/attendeesController"
	"wscmakebygo.com/internal/controller/eventsController"
)

func hookEventsRoute(api *echo.Group) {
	api.GET("/events", eventsController.GetEvents)
}

func hookEventDetailRoute(api *echo.Group) {
	api.GET("/organizers/:organizerSlug/events/:eventSlug", eventsController.GetEventDetail)
}

func hookLoginRoute(api *echo.Group) {
	api.POST("/login", attendeesController.AttendeesLogin)
}

func hookLogoutRoute(api *echo.Group) {
	api.POST("/logout", attendeesController.AttendeesLogout)
}

func HookRoute() {
	var api = route.GetRoute().Group("/api/v1")
	hookEventsRoute(api)
	hookEventDetailRoute(api)
	hookLoginRoute(api)
	hookLogoutRoute(api)
}

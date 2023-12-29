package internal

import (
	"wscmakebygo.com/global/route"
	"wscmakebygo.com/internal/controller/eventsController"
)

func hookEventsRoute() {
	var api = route.GetRoute().Group("/api/v1")
	api.GET("/events", eventsController.GetEvents)
}

func hookEventDetailRoute() {
	var api = route.GetRoute().Group("/api/v1")
	api.GET("/organizers/:organizerSlug/events/:eventSlug", eventsController.GetEventDetail)
}

func HookRoute() {
	hookEventsRoute()
	hookEventDetailRoute()
}

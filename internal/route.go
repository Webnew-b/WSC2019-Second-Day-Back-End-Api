package internal

import (
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/controller/eventsController"
	"wscmakebygo.com/internal/controller/userController"
)

func hookUserRoute() {
	global.Router.GET("/", userController.GetUser)
}

func hookEventsRoute() {
	var api = global.Router.Group("/api/v1")
	api.GET("/events", eventsController.GetEvents)
}

func hookEventDetailRoute() {
	var api = global.Router.Group("/api/v1")
	api.GET("/organizers/:organizerSlug/events/:eventSlug", eventsController.GetEventDetail)
}

func HookRoute() {
	hookUserRoute()
	hookEventsRoute()
	hookEventDetailRoute()
}

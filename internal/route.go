package internal

import (
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/controller/eventsController"
	"wscmakebygo.com/internal/controller/userController"
)

func hookUserRoute() {
	global.Router.GET("/", userController.GetUser)
}

func hookEvensRoute() {
	var api = global.Router.Group("/api/v1")
	api.GET("/events", eventsController.GetEvents)
}

func HookRoute() {
	hookUserRoute()
	hookEvensRoute()
}

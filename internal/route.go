package internal

import (
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/controller/userController"
)

func hookUserRoute() {
	global.Router.GET("/", userController.GetUser)
}

func HookRoute() {
	hookUserRoute()
}

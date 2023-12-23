package start

import (
	"github.com/labstack/echo/v4"
	"log"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal"
)

func createHttpServer() {
	log.Println("starting http Server")
	global.Router = echo.New()
	internal.HookRoute()
	global.Router.Logger.Fatal(global.Router.Start(createServerAddr()))
}

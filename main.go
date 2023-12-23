package main

import (
	"wscmakebygo.com/start"
	"wscmakebygo.com/tools"
)

func main() {
	start.Init()
	start.StopServe()
	tools.Close()
}

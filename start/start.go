package start

import "wscmakebygo.com/tools"

func Init() {
	tools.Log.Println("Server Start")
	createConfig()
	crateDbConnect()
	createHttpServer()
}

func StopServe() {
	tools.Log.Println("stopping Serve")
}

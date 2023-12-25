package start

import "wscmakebygo.com/tools"

func Init() {
	tools.Log.Println("Server is Starting")
	createConfig()
	crateDbConnect()
	crateRedisConnect()
	createHttpServer()
	tools.Log.Println("Server is Started")
}

func StopServe() {
	tools.Log.Println("stopping Server")
}

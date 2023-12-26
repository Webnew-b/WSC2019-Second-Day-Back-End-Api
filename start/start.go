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

func StartDbConnect() {
	createConfig()
	crateDbConnect()
}

func StopServe() {
	tools.Log.Println("stopping Server")
	// todo 停止日志读写以及清退存在的协程
}

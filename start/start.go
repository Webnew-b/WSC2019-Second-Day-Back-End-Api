package start

import (
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/global/envConfig"
	"wscmakebygo.com/global/redisConn"
	"wscmakebygo.com/global/route"
	"wscmakebygo.com/internal"
	"wscmakebygo.com/tools"
)

func Init() {
	tools.Log.Println("Server is Starting")
	envConfig.InitVal()
	database.InitVal()
	redisConn.InitVal()
	route.InitVal()
	internal.HookRoute()
	route.StartRoute()
	tools.Log.Println("Server is Started")
}

func StartDbConnect() {
	envConfig.InitVal()
	database.InitVal()
}

func StopServe() {
	tools.Log.Println("stopping Server")
	// todo 停止日志读写以及清退存在的协程
}

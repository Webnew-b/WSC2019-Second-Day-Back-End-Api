package main

import (
	"wscmakebygo.com/start"
	"wscmakebygo.com/tools"
)

func main() {
	start.Init()
	start.StopServe()
	// todo http服务go化，主进程阻塞，监听程序停止信号。
	tools.Close()
}

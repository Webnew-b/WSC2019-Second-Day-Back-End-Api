package start

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"wscmakebygo.com/global"
	"wscmakebygo.com/tools"
)

func crateRedisAddr() string {
	Addr := fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port)
	return Addr
}

func crateRedisConnect() {
	tools.Log.Println("Connecting Redis")
	logStr := crateRedisAddr()
	rdb := redis.NewClient(&redis.Options{
		Addr:     logStr,
		Password: global.Config.Redis.Password, // no password set
		DB:       global.Config.Redis.DB,       // use default DB
	})
	global.Rdb = rdb
	tools.Log.Println("Connect Redis success")
}

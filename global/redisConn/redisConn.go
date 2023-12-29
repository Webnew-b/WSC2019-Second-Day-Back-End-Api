package redisConn

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"wscmakebygo.com/config"
	"wscmakebygo.com/global/envConfig"
	"wscmakebygo.com/tools"
)

var (
	rdb         *redis.Client
	once        sync.Once
	redisConfig *config.Redis
)

func GetRedis() *redis.Client {
	if rdb == nil {
		panic("redis not initialized")
	}
	return rdb
}

func crateRedisAddr() string {
	Addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	return Addr
}

func InitVal() {
	redisConfig = envConfig.GetConfig().Redis
	once.Do(func() {
		tools.Log.Println("Connecting Redis")
		logStr := crateRedisAddr()
		init := redis.NewClient(&redis.Options{
			Addr:     logStr,
			Password: redisConfig.Password, // no password set
			DB:       redisConfig.DB,       // use default DB
		})
		tools.Log.Println("created Redis connection:" + logStr)
		rdb = init
		tools.Log.Println("Connect Redis success")
	})
}

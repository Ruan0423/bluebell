package redis

import (
	"fmt"
	"web_framework/settings"

	"github.com/go-redis/redis/v8"

)

var Rdb *redis.Client
func Init(cfg *settings.Redisconfig)(err error){
	adrr := fmt.Sprint("%s:%d",cfg.Host,cfg.Port)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     adrr, // Redis 服务器地址
		Password: "",               // 如果有密码的话
		DB:       0,                // 默认数据库是 0
	})
	return
}
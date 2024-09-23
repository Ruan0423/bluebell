package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Rdb *redis.Client
func Init()(err error){
	adrr := fmt.Sprint("%s:%s",viper.GetString("redis.host"),viper.GetString("redis.port"))
	Rdb = redis.NewClient(&redis.Options{
		Addr:     adrr, // Redis 服务器地址
		Password: "",               // 如果有密码的话
		DB:       0,                // 默认数据库是 0
	})
	return
}
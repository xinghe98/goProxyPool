package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// 直接声明一个全局的数据库操作对象
var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	passwd := viper.GetString("DB.passwd")
	url := viper.GetString("DB.url")

	Rdb = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: passwd,
		DB:       0,
	})
	// 测试连接
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("redis连接成功") // 输出: PONG
}

package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Redis struct {
	Rdb *redis.Client
	Ctx context.Context
}

func NewRedisClient() *Redis {
	rdb, ctx := initRedis()
	return &Redis{
		Rdb: rdb,
		Ctx: ctx,
	}
}

func initRedis() (*redis.Client, context.Context) {
	passwd := viper.GetString("DB.passwd")
	url := viper.GetString("DB.url")
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: passwd,
		DB:       0,
	})
	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("redis连接成功") // 输出: PONG
	return rdb, ctx
}

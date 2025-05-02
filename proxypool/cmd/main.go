package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/xinghe98/goProxyPool/common"
	"github.com/xinghe98/goProxyPool/common/db"
	fetchmoudle "github.com/xinghe98/goProxyPool/pkg/fetchMoudle"
)

func main() {
	common.InitConfig()
	db.InitRedis()
	url := viper.GetString("ipurl")
	// 此处添加其他代理商获取ip
	getip := fetchmoudle.NewFetch(url)
	ipgeters := []common.IPGeter{
		&getip,
	}
	for _, v := range ipgeters {
		ips := v.GetIps()
		for _, ip := range ips {
			db.Rdb.ZAdd(db.Ctx, "zset", redis.Z{Score: 100, Member: ip})
		}
		fmt.Println(ips)
	}
}

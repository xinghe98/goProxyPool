package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/xinghe98/goProxyPool/common"
	"github.com/xinghe98/goProxyPool/pkg/fetchMoudle"
	"github.com/xinghe98/goProxyPool/pkg/storgeMoudle"
)

func main() {
	common.InitConfig()
	url := viper.GetString("ipurl")
	// 初始化储存器（同时连接Redis）
	storge := storgeMoudle.NewStorge()

	// 此处添加其他代理商获取ip
	getip := fetchMoudle.NewFetch(url)
	ipgeters := []common.IPGeter{
		&getip,
	}
	for _, v := range ipgeters {
		go func(v common.IPGeter) {
			ips := v.GetIps()
			for _, ip := range ips {
				storge.SaveIp(ip, 20)
			}
			fmt.Println(ips)
		}(v)
	}
	select {}
}

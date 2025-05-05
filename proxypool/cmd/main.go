package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/xinghe98/goProxyPool/common"
	"github.com/xinghe98/goProxyPool/pkg/detectMoudle"
	"github.com/xinghe98/goProxyPool/pkg/fetchMoudle"
	"github.com/xinghe98/goProxyPool/pkg/storgeMoudle"
)

func main() {
	common.InitConfig()
	url := viper.GetString("ipurl")

	// 初始化储存器（同时连接Redis）
	storge := storgeMoudle.NewStorge()

	// 初始化检测器
	detecter := detectMoudle.NewDetect(storge)

	// 此处添加其他代理商获取ip
	// getip := fetchMoudle.NewFetch(url)
	getip := fetchMoudle.NewFetchTest(url) // 测试
	ipgeters := []common.IPGeter{
		&getip,
	}

	// 获取器
	for _, v := range ipgeters {
		go func(v common.IPGeter) {
			ips := v.GetIps()
			for _, ip := range ips {
				storge.SaveIp(ip, 30)
			}
			log.Println("新一批ip保存成功")
		}(v)
	}

	// 检测器:每轮取50个检测
	go func(detecter *detectMoudle.DetectIP) {
		// 取总ip数
		count := storge.GetCount()
		res := storge.GetSomeIp(0, int64(count))
		detecter.TestIp(res)
	}(detecter)

	select {}
}

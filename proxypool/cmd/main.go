package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/xinghe98/goProxyPool/common"
	"github.com/xinghe98/goProxyPool/pkg/apiMoudle"
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

	// 初始化获取器
	// 此处添加其他代理商获取ip
	getip := fetchMoudle.NewFetch(url)
	// getip := fetchMoudle.NewFetchTest(url) // 测试
	ipgeters := []common.IPGeter{
		&getip,
	}
	// 启动一次获取器
	fetchMoudle.Run(ipgeters, storge)
	// 启动一次检测器
	detectMoudle.Run(detecter, storge)

	// 初始化httpapi
	api := apiMoudle.New("3000", storge)

	// 启动gin服务器
	go api.Run()

	// 每3分钟做一次获取
	tickerFetch := time.NewTicker(180 * time.Second)
	// 每2分钟做一次测试
	tickerTest := time.NewTicker(120 * time.Second)

	for {
		select {

		case <-tickerTest.C:
			fmt.Println("执行检测任务，当前时间:", time.Now().Format("2006-01-02 15:04:05"))
			// 在这里执行你的任务
			detectMoudle.Run(detecter, storge)

		case <-tickerFetch.C:
			fmt.Println("执行获取ip任务，当前时间:", time.Now().Format("2006-01-02 15:04:05"))
			// 在这里执行你的任务
			fetchMoudle.Run(ipgeters, storge)
		}
	}
}

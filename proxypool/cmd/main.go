package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/xinghe98/goProxyPool/common"
)

func main() {
	common.InitConfig()
	url := viper.GetString("DB.url")
	fmt.Println(url)
}

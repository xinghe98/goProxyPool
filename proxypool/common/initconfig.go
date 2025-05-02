package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

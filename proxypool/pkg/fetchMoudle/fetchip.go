package fetchMoudle

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/xinghe98/goProxyPool/common"
)

// 获取ip的方式,若是新增方式需要实现IPGeter接口
type fetchIp struct {
	Url string
}

func NewFetch(url string) fetchIp {
	return fetchIp{
		Url: url,
	}
}

func (f *fetchIp) GetIps() []string {
	resp, err := http.Get(f.Url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return strings.Split(string(body), "\n")
}

// 获取器运行方法
func Run(ipgeters []common.IPGeter, storge common.IPStorger) {
	// 初始化储存器（同时连接Redis）
	for _, v := range ipgeters {
		go func(v common.IPGeter) {
			ips := v.GetIps()
			for _, ip := range ips {
				storge.SaveIp(ip, 30)
			}
			log.Println("新一批ip保存成功")
		}(v)
	}
}

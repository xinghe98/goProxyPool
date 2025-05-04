package detectMoudle

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/xinghe98/goProxyPool/common"
)

type DetectIP struct {
	ipstorge common.IPStorger
}

// 初始化校验器
func NewDetect(storge common.IPStorger) *DetectIP {
	return &DetectIP{
		ipstorge: storge,
	}
}

// 并发测试ip可用性
func (d *DetectIP) TestIp(ip []string) {
	workerCount := runtime.NumCPU()
	var wg sync.WaitGroup
	taskChan := make(chan string, len(ip))
	// 1. 启动worker检测器
	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d.checkWorker(taskChan)
		}()
	}

	// 2. 发送任务到channel
	for _, i := range ip {
		taskChan <- i
	}
	close(taskChan)

	// 3. 阻塞主线程
	go func() {
		wg.Wait()
	}()
}

func (d *DetectIP) checkWorker(taskChan <-chan string) {
	for ip := range taskChan {
		res, _ := d.requestCheckUrl(ip)
		if !res {
			log.Printf("❌IP: %s,不可用", ip)
			d.ipstorge.SaveIp(ip, -10)
		} else {
			d.ipstorge.SaveIp(ip, 100)
		}
	}
}

func (d *DetectIP) requestCheckUrl(proxyip string) (bool, error) {
	testurl := viper.GetString("checkurl")
	ip, _ := url.Parse(fmt.Sprintf("http://%s", proxyip))
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(ip), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("GET", testurl, nil)
	res, err := client.Do(req)
	if err != nil {
		return false, errors.New("访问错误")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	log.Printf("✅ IP：%s,检测返回：%s", proxyip, string(body))
	return true, nil
}

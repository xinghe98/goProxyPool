package apiMoudle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xinghe98/goProxyPool/common"
)

type getApi struct {
	RunAddr  string
	ipstorge common.IPStorger
}

func New(addr string, storge common.IPStorger) *getApi {
	runaddr := fmt.Sprintf(":%s", addr)
	return &getApi{
		RunAddr:  runaddr,
		ipstorge: storge,
	}
}

func (g *getApi) Run() {
	r := gin.Default()
	// 测试接口
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 获取一个ip
	r.GET("", func(ctx *gin.Context) {
		ip, err := g.ipstorge.GetIp()
		if err != nil {
			ctx.JSON(http.StatusNotExtended, gin.H{
				"error": "暂无ip可以使用",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"result": ip,
		})
	})

	r.Run(g.RunAddr)
}

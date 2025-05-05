package apiMoudle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getApi struct {
	RunAddr string
}

func New(addr string) *getApi {
	runaddr := fmt.Sprintf(":%s", addr)
	return &getApi{
		RunAddr: runaddr,
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

	r.Run(g.RunAddr)
}

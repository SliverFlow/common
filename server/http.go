package server

import (
	"fmt"
	"github.com/SliverFlow/core/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Http struct {
	Server *http.Server
}

type ApiGroup interface {
	InitApi(r *gin.Engine)
}

// NewHttp 创建 http 服务
func NewHttp(logger *zap.Logger, c *config.HttpServer, api ApiGroup) *Http {
	R := gin.Default()
	api.InitApi(R)
	time.Sleep(500 * time.Millisecond)
	logger.Info(fmt.Sprintf("[项目运行于]：http://127.0.0.1:%d", c.Port))
	return &Http{
		Server: &http.Server{
			Addr:           fmt.Sprintf(":%d", c.Port),
			Handler:        R,
			ReadTimeout:    20 * time.Second,
			WriteTimeout:   20 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

// ListenServer 监听服务
func (h *Http) ListenServer() error {
	return h.Server.ListenAndServe()
}

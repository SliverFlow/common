package server

import (
	"context"
	"fmt"
	"github.com/SliverFlow/core/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	var err error
	go func() {
		if err = h.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			_ = h.Server.Close()
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down the server...")
	if err := h.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
	fmt.Println("Server exited successfully")
	return nil
}

// Shutdown 关闭服务
func (h *Http) Shutdown(ctx context.Context) error {
	return h.Server.Shutdown(ctx)
}

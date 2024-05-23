package middleware

import (
	"github.com/SliverFlow/core/config"
	gtimeout "github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Timeout struct {
	config *config.HttpServer
}

// NewTimeout 创建超时中间件
func NewTimeout(logger *zap.Logger, c *config.HttpServer) *Timeout {
	return &Timeout{
		config: c,
	}
}

func (t *Timeout) Handle() gin.HandlerFunc {
	tot := t.config.Timeout
	if tot <= 2 || tot >= 30 {
		tot = 5
	}
	return gtimeout.New(
		gtimeout.WithTimeout(time.Duration(tot)*time.Second),
		gtimeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		gtimeout.WithResponse(func(c *gin.Context) {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"code": http.StatusRequestTimeout,
				"msg":  "请求超时",
			})
		}),
	)
}

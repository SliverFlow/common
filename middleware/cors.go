package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cors struct{}

// NewCors 创建跨域中间件
func NewCors() *Cors {
	return &Cors{}
}

func (cors *Cors) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, new-token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.Status(http.StatusNoContent)
		}
		c.Next()
	}
}

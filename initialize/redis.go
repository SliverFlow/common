package initialize

import (
	"context"
	"fmt"
	"github.com/SliverFlow/core/config"
)

// Redis 初始化 Redis 连接
func Redis(c *config.Redis) *redis.Client {

	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.DB,
	})

	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
		return nil
	}

	return cli
}

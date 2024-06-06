package initialize

import (
	"context"
	"fmt"
	"github.com/SliverFlow/core/config"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// Redis 初始化 Redis 连接
func Redis(c *config.Redis) (*redis.Client, error) {

	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.DB,
	})

	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return cli, nil
}

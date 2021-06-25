package com

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/dsl/user/internal/cfg"
)

var Redis *redis.Client

func initRedis() error {
	utils.Assert(cfg.Instance != nil, "配置文件未加载")

	Redis = redis.NewClient(&redis.Options{
		Addr:     cfg.Instance.Redis.Hosts[0],
		Username: cfg.Instance.Redis.User,
		Password: cfg.Instance.Redis.Pass,
	})
	return Redis.Ping(context.TODO()).Err()
}

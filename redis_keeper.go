package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const TTL = 0

type RedisKeeper struct {
	cn  redis.Client
	ctx context.Context
}

func (k RedisKeeper) Get(key string) (string, error) {
	return k.cn.Get(k.ctx, key).Result()
}

func (k RedisKeeper) Set(key string, message string) error {
	return k.cn.Set(k.ctx, key, message, TTL).Err()
}

func (k RedisKeeper) Clean(key string) error {
	return k.cn.Del(k.ctx, key).Err()
}

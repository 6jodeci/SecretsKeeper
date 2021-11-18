package main

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

const TTL = 0

type RedisKeeper struct {
	cn  redis.Client
	ctx context.Context
}

func (k RedisKeeper) Get(key string) (string, error) {
	val, err := k.cn.Get(k.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New(NOT_FOUND_ERROR)
	}
	return val, err
}

func (k RedisKeeper) Set(key string, message string) error {
	return k.cn.Set(k.ctx, key, message, TTL).Err()
}

func (k RedisKeeper) Clean(key string) error {
	return k.cn.Del(k.ctx, key).Err()
}

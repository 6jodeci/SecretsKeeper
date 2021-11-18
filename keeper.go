package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

const NOT_FOUND_ERROR = "not_found"

type Keeper interface {
	Get(key string) (string, error)
	Set(key string, message string) error
	Clean(key string) error
}
type DummyKeeper struct {
	memory map[string]string
}

func (k *DummyKeeper) Get(key string) (string, error) {
	value, ok := k.memory[key]
	if !ok {
		return "", errors.New(NOT_FOUND_ERROR)
	}
	return value, nil
}

func (k *DummyKeeper) Set(key string, message string) error {
	k.memory[key] = message
	return nil
}

func (k *DummyKeeper) Clean(key string) error {
	delete(k.memory, key)
	return nil
}

func getDummyKeeper() Keeper {
	return &DummyKeeper{make(map[string]string)}
}

func getRedisKeeper() Keeper {
	return RedisKeeper{*redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}), context.Background()}
}

package main

import "errors"

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
		return "", errors.New("message not found")
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

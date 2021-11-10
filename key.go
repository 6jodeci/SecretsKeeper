package main

const DUMMY_TEST_KEY = "test_key"

type KeyGenerator interface {
	Get() string
}

type DummyKeyBuilder struct {
}

func (k *DummyKeyBuilder) Get() string {
	return DUMMY_TEST_KEY
}

func getKeyBuilder() KeyGenerator {
	return &DummyKeyBuilder{}
}

var keyBuilder = getKeyBuilder()

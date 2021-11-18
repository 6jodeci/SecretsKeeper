package main

import "testing"

func TestDummyKeyGenerator(t *testing.T) {
	dummyKeyBuilder := DummyKeyBuilder{}
	key, _ := dummyKeyBuilder.Get()
	if key != DUMMY_TEST_KEY {
		t.Error("ERROR: bad dummy key")
	}
}

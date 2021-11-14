package main

import "testing"

func TestDummyKeyGenerator(t *testing.T) {
	dummyKeyBuilder := DummyKeyBuilder{}
	if dummyKeyBuilder.Get() != DUMMY_TEST_KEY {
		t.Error("bad dummy key")
	}
}

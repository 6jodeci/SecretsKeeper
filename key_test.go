package main

import "testing"

func TestDummyKeyGenerator(t *testing.T) {
	dummy_key_builder := DummyKeyBuilder{}
	if dummy_key_builder.Get() != DUMMY_TEST_KEY {
		t.Error("bad dummy key")
	}
}

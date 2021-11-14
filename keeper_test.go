package main

import "testing"

func TestDummyKeeperSet(t *testing.T) {
	keeper := DummyKeeper{make(map[string]string)}
	key := "foo"
	value := "bar"
	keeper.Set(key, value)
	if keeper.memory[key] != value {
		t.Error("ERROR: bad memory value")
	}
}

func TestDummyKeeperGet(t *testing.T) {
	keeper := DummyKeeper{make(map[string]string)}
	key := "foo"
	value := "bar"
	keeper.memory[key] = value
	valueFromGet, _ := keeper.Get(key)
	if valueFromGet != value {
		t.Error("ERROR: bad value from get")
	}
}

func TestDummyKeeperClean(t *testing.T) {
	keeper := DummyKeeper{make(map[string]string)}
	key := "foo"
	value := "bar"
	keeper.memory[key] = value
	keeper.Clean(key)
	_, ok := keeper.memory[key]
	if ok {
		t.Error("ERROR: clean does not work")
	}
}

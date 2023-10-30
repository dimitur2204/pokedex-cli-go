package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/dimitur2204/pokedex-cli-go/internal/pokeapi/internal/pokecache"
)

func TestAddGet(t *testing.T) {

	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		mux := &sync.Mutex{}
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Set(c.key, &c.val, mux)
			val, ok := cache.Get(c.key, mux)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(*val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	mux := &sync.Mutex{}
	data := []byte("testdata")
	cache.Set("https://example.com", &data, mux)

	_, ok := cache.Get("https://example.com", mux)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com", mux)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

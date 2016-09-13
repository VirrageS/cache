package cache

import (
	"testing"
	"time"
)

func Test_CacheSetGet(t *testing.T) {
	cache := NewCache(time.Minute * 10)

	data, exists := cache.Get("hello")
	if exists {
		t.Errorf("Expected empty cache to return no data")
	}

	cache.Set("hello", "world")
	data, exists = cache.Get("hello")
	if !exists {
		t.Errorf("Expected cache to return data for `hello`")
	}

	if data.(string) != "world" {
		t.Errorf("Expected cache to return `world` for `hello`")
	}
}

func Test_CacheCleaning(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cache.Set("hello", "world")
	_, exists := cache.Get("hello")
	if !exists {
		t.Errorf("Expected cache to return data for `hello`")
	}

	time.Sleep(time.Millisecond + time.Millisecond)
	_, exists = cache.Get("hello")
	if exists {
		t.Errorf("Expected cache to clean all data and return nothing for `hello`")
	}
}

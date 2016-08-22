package cache

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
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

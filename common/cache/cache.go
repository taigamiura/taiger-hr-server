package cache

import (
	"sync"
	"time"
)

// SimpleCache は簡易キャッシュ実装です
type SimpleCache struct {
	mu    sync.RWMutex
	items map[string]CacheItem
}

// CacheItem はキャッシュアイテムです
type CacheItem struct {
	Value      interface{}
	Expiration int64
}

// NewCache は新しいキャッシュを作成します
func NewCache() *SimpleCache {
	return &SimpleCache{
		items: make(map[string]CacheItem),
	}
}

// Set はキャッシュにアイテムを設定します
func (c *SimpleCache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}

// Get はキャッシュからアイテムを取得します
func (c *SimpleCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found || time.Now().UnixNano() > item.Expiration {
		return nil, false
	}
	return item.Value, true
}

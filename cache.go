package wuiilcache

import (
	"github.com/wuiilcache/policy"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *policy.Policy
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		// Lazy Initialization
		c.lru = policy.NewCache(policy.LRUNew(c.cacheBytes, nil))
	}
	(*c.lru).Set(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := (*c.lru).Get(key); ok {
		return v.(ByteView), ok
	}

	return
}

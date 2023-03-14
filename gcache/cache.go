package gcache

import "sync"

type cache struct {
	mu            sync.Mutex
	lru           *LRU
	maxCacheBytes int64
}

func (c *cache) add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = NewLRU(c.maxCacheBytes, nil)
	}
	c.lru.Add(key, val)
}

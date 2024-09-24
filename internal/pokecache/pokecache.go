package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	pokeVault          map[string]vaultContent
	mu                 sync.Mutex
	expirationInterval time.Duration
	deleteCache        chan bool
}

type vaultContent struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		expirationInterval: interval,
		pokeVault:          make(map[string]vaultContent),
	}

	go func() {
		c.reapLoop(interval)
	}()

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.pokeVault[key] = vaultContent{createdAt: time.Now(), val: val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	val, ok := c.pokeVault[key]
	c.mu.Unlock()
	if !ok {
		return []byte{}, false
	}
	return val.val, true
}

func (c *Cache) DeleteCache() {
	c.deleteCache <- true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for key, content := range c.pokeVault {
				if time.Since(content.createdAt) > interval {
					c.mu.Lock()
					delete(c.pokeVault, key)
					c.mu.Unlock()
				}
			}
		}
	}
}

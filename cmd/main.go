package main

import "sync"

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string)
}

type CacheImpl struct {
	mu    sync.RWMutex
	cache map[string]string
}

func (c *CacheImpl) Get(key string) (value string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok = c.cache[key]

	return
}

func (c *CacheImpl) Set(key, value string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.cache[key] = value
}

func main() {

}

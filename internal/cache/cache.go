package cache

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Cache --with-expecter=true

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string)
	getHash(key string) int
}

type CacheImpl struct {
	n      int
	caches []*ShardedCacheImpl
}

func NewCache(n int) *CacheImpl {
	var caches = make([]*ShardedCacheImpl, 0, 10)
	for i := 0; i < n; i++ {
		var c = &ShardedCacheImpl{
			Cache: make(map[string]string),
			Mu:    &sync.RWMutex{},
		}

		caches = append(caches, c)
	}

	return &CacheImpl{
		n:      n,
		caches: caches,
	}
}

func (c *CacheImpl) getHash(key string) int {
	hash := sha1.Sum([]byte(key))
	lowerBits := int(hash[0] & ((1 << c.n) - 1))

	return lowerBits % c.n
}

func (c *CacheImpl) Get(key string) (string, bool) {
	i := c.getHash(key)
	if i >= len(c.caches) {
		return "", false
	}

	return c.caches[i].Get(key)
}

func (c *CacheImpl) Set(key, value string) error {
	i := c.getHash(key)
	if i >= len(c.caches) {
		return fmt.Errorf("index out of range")
	}

	c.caches[i].Set(key, value)

	return nil
}

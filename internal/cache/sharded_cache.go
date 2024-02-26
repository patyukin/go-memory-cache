package cache

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=ShardedCache --with-expecter=true --inpackage

import "sync"

type ShardedCache interface {
	get(key string) (string, error)
	set(key, value string)
}

type ShardedCacheImpl struct {
	Mu    *sync.RWMutex
	Cache map[string]string
}

func (sc *ShardedCacheImpl) Get(key string) (value string, ok bool) {
	sc.Mu.RLock()
	defer sc.Mu.RUnlock()

	value, ok = sc.Cache[key]

	return
}

func (sc *ShardedCacheImpl) Set(key, value string) {
	sc.Mu.Lock()
	defer sc.Mu.Unlock()
	sc.Cache[key] = value
}

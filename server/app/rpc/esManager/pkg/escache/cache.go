package escache

import (
	"sync"
)

type CacheClient struct {
	client *sync.Map
}

var (
	cache *CacheClient
	once  sync.Once
)

func NewCache() *CacheClient {
	once.Do(func() {
		cache = &CacheClient{client: new(sync.Map)}
	})
	return cache
}

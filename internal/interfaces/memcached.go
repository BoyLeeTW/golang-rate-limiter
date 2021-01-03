package interfaces

import "github.com/bradfitz/gomemcache/memcache"

type MemcachedClientInterface interface {
	Add(item *memcache.Item) error
	Increment(key string, delta uint64) (newValue uint64, err error)
}

type MemcachedServiceInterface interface {
	IncrementOrAdd(key string, delta uint64, expiration int32) (uint64, error)
}

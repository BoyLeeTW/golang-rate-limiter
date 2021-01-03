package pkg

import (
	"golang-rate-limiter/internal/interfaces"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcachedService struct {
	client interfaces.MemcachedClientInterface
}

func NewMemcachedService(client interfaces.MemcachedClientInterface) *MemcachedService {
	return &MemcachedService{
		client: client,
	}
}

func (s *MemcachedService) IncrementOrAdd(key string, value uint64, expiration int32) (uint64, error) {
	newValue, err := s.client.Increment(key, value)

	if err != nil {
		switch err {
		case memcache.ErrCacheMiss:
			// key doesn't exist
			return s.addOrIncrement(key, value, expiration)
		default:
			// other error
			return 0, err
		}
	}
	return newValue, nil
}

// addOrIncrement use atomically Add first, received ErrNotStored if the key is already created by other request, so it calls Increment if received ErrNotStored. Return (0, err) with all other error.
func (s *MemcachedService) addOrIncrement(key string, value uint64, expiration int32) (uint64, error) {
	err := s.client.Add(&memcache.Item{Key: key, Value: []byte("1"), Expiration: expiration})
	if err != nil {
		switch err {
		case memcache.ErrNotStored:
			// key already exists, increment the value
			newValue, err := s.client.Increment(key, value)
			if err != nil {
				// received other error
				return 0, err
			}
			// increment success
			return newValue, nil
		default:
			// received other error
			return 0, err
		}
	}
	// add success
	return 1, nil
}

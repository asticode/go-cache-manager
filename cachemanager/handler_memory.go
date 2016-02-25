package cachemanager

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// NewHandlerMemory creates a memory handler
func NewHandlerMemory(cleanupInterval time.Duration, prefix string, ttl time.Duration) Handler {
	return &handlerMemory{
		client: cache.New(1*time.Nanosecond, cleanupInterval),
		handler: handler{
			prefix: prefix,
			ttl:    ttl,
		},
	}
}

// NewHandlerMemoryFromConfiguration creates a memory handler based on a configuration
func NewHandlerMemoryFromConfiguration(c ConfigurationMemory) Handler {
	return NewHandlerMemory(
		time.Duration(c.CleanupInterval)*time.Nanosecond,
		c.Prefix,
		time.Duration(c.TTL)*time.Nanosecond,
	)
}

type handlerMemory struct {
	client *cache.Cache
	handler
}

func (h handlerMemory) Decrement(key string, delta uint64) (uint64, error) {
	return h.client.DecrementUint64(h.buildKey(key), delta)
}

func (h handlerMemory) Del(key string) error {
	// Delete
	h.client.Delete(h.buildKey(key))

	// Return
	return nil
}

func (h handlerMemory) Get(key string) (interface{}, error) {
	// Get value
	v, ok := h.client.Get(h.buildKey(key))

	// Return
	if ok {
		return v, nil
	} else {
		return nil, ErrCacheMiss
	}
}

func (h handlerMemory) Increment(key string, delta uint64) (uint64, error) {

	return h.client.IncrementUint64(h.buildKey(key), delta)
}

func (h handlerMemory) Set(key string, value interface{}, ttl time.Duration) error {
	// Set
	h.client.Set(h.buildKey(key), value, h.buildTTL(ttl))

	// Return
	return nil
}

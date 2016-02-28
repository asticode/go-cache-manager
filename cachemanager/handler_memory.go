package cachemanager

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// NewHandlerMemory creates a memory handler
func NewHandlerMemory(cleanupInterval time.Duration, maxSize int64, prefix string, ttl time.Duration) Handler {
	return &handlerMemory{
		client: cache.New(1*time.Nanosecond, cleanupInterval),
		handler: handler{
			prefix: prefix,
			ttl:    ttl,
		},
		maxSize: maxSize,
	}
}

// NewHandlerMemoryFromConfiguration creates a memory handler based on a configuration
func NewHandlerMemoryFromConfiguration(c ConfigurationMemory) Handler {
	return NewHandlerMemory(
		time.Duration(c.CleanupInterval)*time.Nanosecond,
		c.MaxSize,
		c.Prefix,
		time.Duration(c.TTL)*time.Nanosecond,
	)
}

type handlerMemory struct {
	client *cache.Cache
	handler
	maxSize int64
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
	// Initialize
	var i interface{}

	// Get value
	i, ok := h.client.Get(h.buildKey(key))

	// Cache miss
	if !ok {
		return i, ErrCacheMiss
	}
	return i, nil
}

func (h handlerMemory) Increment(key string, delta uint64) (uint64, error) {

	return h.client.IncrementUint64(h.buildKey(key), delta)
}

func (h handlerMemory) Set(key string, value interface{}, ttl time.Duration) error {
	// Check max size
	if int64(h.client.ItemCount()) >= h.maxSize {
		return ErrCacheFull
	}

	// Set
	h.client.Set(h.buildKey(key), value, h.buildTTL(ttl))

	// Return
	return nil
}

func (h handlerMemory) SetOnEvicted(f func(k string, v interface{})) Handler {
	h.client.OnEvicted(f)
	return h
}

package cachemanager

import (
	"strings"
	"github.com/bradfitz/gomemcache/memcache"
)

// NewHandlerMemcache creates a memcache handler
func NewHandlerMemcache(servers string, prefix string, ttl int32) Handler {
	return &handlerMemcache{
		client: memcache.New(strings.Split(servers, ",")...),
		handlerBase: handlerBase{
			prefix: prefix,
			ttl: ttl,
		},
	}
}

// NewHandlerMemcacheFromConfiguration creates a memcache handler based on a configuration
func NewHandlerMemcacheFromConfiguration(c ConfigurationMemcache) Handler {
	return NewHandlerMemcache(c.servers, c.prefix, c.ttl)
}

type handlerMemcache struct {
	client *memcache.Client
	handlerBase
}

func (h handlerMemcache) Decrement(key string, delta uint64) (uint64, error) {
	return h.client.Decrement(h.buildKey(key), delta)
}

func (h handlerMemcache) Del(key string) error {
	return h.client.Delete(h.buildKey(key))
}

func (h handlerMemcache) Get(key string) ([]byte, error) {
	// Initialize
	var v []byte

	// Get item
	i, e := h.client.Get(h.buildKey(key))
	if e == nil {
		v = i.Value
	}

	// Return
	return v, e
}

func (h handlerMemcache) Increment(key string, delta uint64) (uint64, error) {
	return h.client.Increment(h.buildKey(key), delta)
}

func (h handlerMemcache) Set(key string, value []byte, ttl int32) error {
	return h.client.Set(&memcache.Item{
		Key:        h.buildKey(key),
		Value:      value,
		Expiration: h.buildTTL(ttl),
	})
}

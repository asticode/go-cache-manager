package cachemanager

import (
	"strings"

	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

// NewHandlerMemcache creates a memcache handler
func NewHandlerMemcache(servers string, prefix string, ttl time.Duration) Handler {
	return &handlerMemcache{
		client: memcache.New(strings.Split(servers, ",")...),
		handler: handler{
			prefix: prefix,
			ttl:    ttl,
		},
	}
}

// NewHandlerMemcacheFromConfiguration creates a memcache handler based on a configuration
func NewHandlerMemcacheFromConfiguration(c ConfigurationMemcache) Handler {
	return NewHandlerMemcache(c.Servers, c.Prefix, time.Duration(c.TTL)*time.Nanosecond)
}

type handlerMemcache struct {
	client *memcache.Client
	handler
}

func (h handlerMemcache) Decrement(key string, delta uint64) (uint64, error) {
	return h.client.Decrement(h.buildKey(key), delta)
}

func (h handlerMemcache) Del(key string) error {
	return h.client.Delete(h.buildKey(key))
}

func (h handlerMemcache) Get(key string) (interface{}, error) {
	// Initialize
	var v interface{}
	var e error

	// Get item
	i, e := h.client.Get(h.buildKey(key))
	if e != nil {
		return v, e
	}

	// Unserialize
	e = h.unserialize(i.Value, &v)

	// Return
	return v, e
}

func (h handlerMemcache) Increment(key string, delta uint64) (uint64, error) {
	return h.client.Increment(h.buildKey(key), delta)
}

func (h handlerMemcache) Set(key string, value interface{}, ttl time.Duration) error {
	// Initialize
	var e error

	// Serialize
	v, e := h.serialize(value)
	if e != nil {
		return e
	}

	// Return
	return h.client.Set(&memcache.Item{
		Key:        h.buildKey(key),
		Value:      v,
		Expiration: int32(h.buildTTL(ttl).Seconds()),
	})
}

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
	o, e := h.client.Decrement(h.buildKey(key), delta)
	if e != nil && e.Error() == memcache.ErrCacheMiss.Error() {
		return o, ErrCacheMiss
	}
	return o, e
}

func (h handlerMemcache) Del(key string) error {
	return h.client.Delete(h.buildKey(key))
}

func (h handlerMemcache) Get(key string, value interface{}) error {
	// Initialize
	var e error

	// Get item
	i, e := h.client.Get(h.buildKey(key))
	if e != nil && e.Error() == memcache.ErrCacheMiss.Error() {
		return ErrCacheMiss
	} else if e != nil {
		return e
	}

	// Unserialize
	if _, ok := value.(*[]byte); !ok {
		e = h.unserialize(i.Value, value)
	} else {
		value = i
	}

	// Return
	return e
}

func (h handlerMemcache) Increment(key string, delta uint64) (uint64, error) {
	o, e := h.client.Increment(h.buildKey(key), delta)
	if e != nil && e.Error() == memcache.ErrCacheMiss.Error() {
		return o, ErrCacheMiss
	}
	return o, e
}

func (h handlerMemcache) Set(key string, value interface{}, ttl time.Duration) error {
	// Initialize
	var e error
	var v []byte

	// Serialize
	if _, ok:= value.([]byte); !ok {
		v, e = h.serialize(value)
		if e != nil {
			return e
		}
	} else {
		v = value.([]byte)
	}

	// Return
	return h.client.Set(&memcache.Item{
		Key:        h.buildKey(key),
		Value:      v,
		Expiration: int32(h.buildTTL(ttl).Seconds()),
	})
}

func (h handlerMemcache) SetOnEvicted(f func (k string, v interface{})) Handler {
	panic("not yet implemented")
}

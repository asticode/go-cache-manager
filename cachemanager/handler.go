package cachemanager

import "time"

type Handler interface {
	Decrement(key string, delta uint64) (uint64, error)
	Del(k string) error
	Get(k string) (interface{}, error)
	Increment(key string, delta uint64) (uint64, error)
	SetOnEvicted(f func(k string, v interface{})) Handler
	Set(k string, v interface{}, ttl time.Duration) error
	Test() error
}

type handler struct {
	prefix string
	ttl    time.Duration
}

func (h handler) buildKey(k string) string {
	return h.prefix + k
}

func (h handler) buildTTL(ttl time.Duration) time.Duration {
	if ttl == time.Duration(-1) {
		return h.ttl
	}
	return ttl
}

func MockHandler() Handler {
	return NewHandlerMemory(
		time.Duration(500)*time.Nanosecond,
		200,
		"mocked_handler:",
		-1,
	)
}

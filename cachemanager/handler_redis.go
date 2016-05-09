package cachemanager

import (
	"gopkg.in/redis.v3"
	"time"
)

// NewHandlerRedis creates a redis handler
func NewHandlerRedis(redisConfig *redis.Options, prefix string, ttl time.Duration) Handler {
	return &handlerRedis{
		client: redis.NewClient(redisConfig),
		handler: handler{
			prefix: prefix,
			ttl:    ttl,
		},
	}
}

type handlerRedis struct {
	client *redis.Client
	handler
}

func (h handlerRedis) Del(key string) error {
	return h.client.Del(h.buildKey(key)).Err()
}

func (h handlerRedis) Get(key string) (interface{}, error) {
	// Initialize
	var o interface{}

	// Get item
	i, e := h.client.Get(h.buildKey(key)).Result()
	if e != nil && e == redis.Nil {
		return o, ErrCacheMiss
	}

	return []byte(i), e
}

func (h handlerRedis) Set(key string, value interface{}, ttl time.Duration) error {
	return h.client.Set(h.buildKey(key), value, ttl).Err()
}

func (h handlerRedis) Increment(key string, delta uint64) (uint64, error) {
	res, err := h.client.IncrBy(h.buildKey(key), int64(delta)).Result()
	if err == redis.Nil {
		return 0, ErrCacheMiss
	} else if err != nil {
		return 0, err
	}

	return uint64(res), err
}

func (h handlerRedis) Decrement(key string, delta uint64) (uint64, error) {
	res, err := h.client.DecrBy(h.buildKey(key), int64(delta)).Result()
	if err == redis.Nil {
		return 0, ErrCacheMiss
	} else if err != nil {
		return 0, err
	}

	return uint64(res), err
}

func (h handlerRedis) SetOnEvicted(f func(k string, v interface{})) Handler {
	panic("not yet implemented")
}

func (h handlerRedis) Test() error {
	panic("not yet implemented")
}

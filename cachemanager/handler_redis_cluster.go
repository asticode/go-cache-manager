package cachemanager

import (
	"gopkg.in/redis.v3"
	"time"
)

// NewHandlerRedis creates a redis handler
func NewHandlerRedisCluster(redisConfig *redis.ClusterOptions, prefix string, ttl time.Duration) Handler {
	return &handlerRedisCluster{
		client: redis.NewClusterClient(redisConfig),
		handler: handler{
			prefix: prefix,
			ttl:    ttl,
		},
	}
}

func NewHandlerRedisClusterFromConfiguration(conf *ConfigurationRedisCluster) Handler {
	return NewHandlerRedisCluster(&redis.ClusterOptions{
		Addrs:         conf.Addrs,
	}, conf.Prefix, time.Duration(conf.TTL))
}

type handlerRedisCluster struct {
	client *redis.ClusterClient
	handler
}

func (h handlerRedisCluster) Del(key string) error {
	return h.client.Del(h.buildKey(key)).Err()
}

func (h handlerRedisCluster) Get(key string) (interface{}, error) {
	// Initialize
	var o interface{}

	// Get item
	i, e := h.client.Get(h.buildKey(key)).Result()
	if e != nil && e == redis.Nil {
		return o, ErrCacheMiss
	}

	return []byte(i), e
}

func (h handlerRedisCluster) Set(key string, value interface{}, ttl time.Duration) error {
	return h.client.Set(h.buildKey(key), value, ttl).Err()
}

func (h handlerRedisCluster) Increment(key string, delta uint64) (uint64, error) {
	res, err := h.client.IncrBy(h.buildKey(key), int64(delta)).Result()
	if err == redis.Nil {
		return 0, ErrCacheMiss
	} else if err != nil {
		return 0, err
	}

	return uint64(res), err
}

func (h handlerRedisCluster) Decrement(key string, delta uint64) (uint64, error) {
	res, err := h.client.DecrBy(h.buildKey(key), int64(delta)).Result()
	if err == redis.Nil {
		return 0, ErrCacheMiss
	} else if err != nil {
		return 0, err
	}

	return uint64(res), err
}

func (h handlerRedisCluster) SetOnEvicted(f func(k string, v interface{})) Handler {
	panic("not yet implemented")
}

func (h handlerRedisCluster) Test() error {
	panic("not yet implemented")
}

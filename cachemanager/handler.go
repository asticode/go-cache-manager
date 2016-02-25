package cachemanager

type Handler interface {
	Decrement(key string, delta uint64) (uint64, error)
	Del(k string) error
	Get(k string) (interface{}, error)
	Increment(key string, delta uint64) (uint64, error)
	Set(k string, v interface{}, ttl int32) error
}

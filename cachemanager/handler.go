package cachemanager

type Handler interface {
	Decode(i []byte, o interface{}) error
	Decrement(key string, delta uint64) (uint64, error)
	Del(k string) error
	Encode(d interface{}) ([]byte, error)
	Get(k string) ([]byte, error)
	Increment(key string, delta uint64) (uint64, error)
	Set(k string, v []byte, ttl int32) error
}



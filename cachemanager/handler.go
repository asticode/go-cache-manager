package cachemanager

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Handler interface {
	Decrement(key string, delta uint64) (uint64, error)
	Del(k string) error
	Get(k string) (interface{}, error)
	Increment(key string, delta uint64) (uint64, error)
	Set(k string, v interface{}, ttl time.Duration) error
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

func (h handler) serialize(d interface{}) ([]byte, error) {
	// Initialize
	var buffer bytes.Buffer

	// Encode
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(d)

	return buffer.Bytes(), err
}

func (h handler) unserialize(i []byte, o interface{}) error {
	// Decode
	dec := gob.NewDecoder(bytes.NewBuffer(i))
	return dec.Decode(o)
}

package cachemanager

import (
	"bytes"
	"encoding/json"
)

type handlerBase struct {
	prefix string
	ttl    int32
}

func (h handlerBase) buildKey(k string) string {
	return h.prefix + k
}

func (h handlerBase) buildTTL(ttl int32) int32 {
	if ttl == -1 {
		return h.ttl
	}
	return ttl
}

func (h handlerBase) Encode(d interface{}) ([]byte, error) {
	// Initialize
	var buffer bytes.Buffer

	// Encode
	enc := json.NewEncoder(&buffer)
	err := enc.Encode(d)

	return buffer.Bytes(), err
}

func (h handlerBase) Decode(i []byte, o interface{}) error {
	// Decode
	dec := json.NewDecoder(bytes.NewBuffer(i))
	return dec.Decode(o)
}
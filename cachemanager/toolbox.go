package cachemanager

import (
	"bytes"
	"encoding/gob"
)

func ToBytes(d interface{}) ([]byte, error) {
	// Initialize
	var err error
	var buffer bytes.Buffer

	// Encode
	if d != nil {
		enc := gob.NewEncoder(&buffer)
		err = enc.Encode(d)
	}

	return buffer.Bytes(), err
}

func FromBytes(i []byte, o interface{}) error {
	// Decode
	dec := gob.NewDecoder(bytes.NewBuffer(i))
	return dec.Decode(o)
}

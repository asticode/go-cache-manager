package cachemanager

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandlerMemory(t *testing.T) {
	// Initialize
	k := "test"
	v := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	c := ConfigurationMemory{
		CleanupInterval: 500,
		Configuration: Configuration{
			Prefix: "test_",
		},
	}
	m := NewHandlerMemoryFromConfiguration(c)

	// Set
	m.Set(k, v, time.Duration(5)*time.Microsecond)
	vc, e := m.Get(k)
	assert.NoError(t, e)
	assert.Equal(t, v, vc)

	// Wait for expiration
	time.Sleep(time.Duration(5) * time.Microsecond)
	vc, e = m.Get(k)
	assert.EqualError(t, e, ErrCacheMiss.Error())

	// Del
	m.Set(k, v, time.Duration(5)*time.Microsecond)
	m.Del(k)
	vc, e = m.Get(k)
	assert.EqualError(t, e, ErrCacheMiss.Error())

	// Increment
	m.Set(k, uint64(5), time.Duration(20)*time.Microsecond)
	_, e = m.Increment(k, 1)
	assert.NoError(t, e)
	vc, e = m.Get(k)
	assert.NoError(t, e)
	assert.Equal(t, uint64(6), vc)

	// Decrement
	m.Decrement(k, 2)
	vc, e = m.Get(k)
	assert.NoError(t, e)
	assert.Equal(t, uint64(4), vc)
}

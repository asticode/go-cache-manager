package cachemanager

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandlerMemory(t *testing.T) {
	// Initialize
	var ic uint64
	k := "test"
	kFail := "test_1"
	v := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	vc := make(map[string]string)
	c := ConfigurationMemory{
		CleanupInterval: 500,
		Configuration: Configuration{
			Prefix: "test_",
		},
		MaxSize: 1,
	}
	m := NewHandlerMemoryFromConfiguration(c)

	// Set
	m.Set(k, v, time.Duration(100)*time.Microsecond)
	e := m.Get(k, &vc)
	assert.NoError(t, e)
	assert.Equal(t, v, vc)

	// Wait for expiration
	time.Sleep(time.Duration(100) * time.Microsecond)
	e = m.Get(k, &vc)
	assert.EqualError(t, e, ErrCacheMiss.Error())

	// Del
	m.Set(k, v, time.Duration(100)*time.Microsecond)
	m.Del(k)
	e = m.Get(k, &vc)
	assert.EqualError(t, e, ErrCacheMiss.Error())

	// Increment
	m.Set(k, uint64(5), time.Duration(100)*time.Microsecond)
	_, e = m.Increment(k, 1)
	assert.NoError(t, e)
	e = m.Get(k, &ic)
	assert.NoError(t, e)
	assert.Equal(t, uint64(6), ic)

	// Decrement
	m.Decrement(k, 2)
	e = m.Get(k, &ic)
	assert.NoError(t, e)
	assert.Equal(t, uint64(4), ic)

	// Max size
	m.Del(k)
	e = m.Set(k, "test", time.Duration(100)*time.Microsecond)
	assert.NoError(t, e)
	e = m.Set(kFail, "test", time.Duration(100)*time.Microsecond)
	assert.EqualError(t, e, ErrCacheFull.Error())
}

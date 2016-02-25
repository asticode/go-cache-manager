// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cachemanager

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransformKey(t *testing.T) {
	// Initialize
	h := handler{
		prefix: "prefix_",
	}

	// Assert
	assert.Equal(t, "prefix_test", h.buildKey("test"))
}

func TestTransformTtl(t *testing.T) {
	// Initialize
	h := handler{
		ttl: time.Duration(5),
	}

	// Assert
	assert.Equal(t, time.Duration(5), h.buildTTL(time.Duration(-1)))
	assert.Equal(t, time.Duration(3), h.buildTTL(time.Duration(3)))
}

func TestSerialize(t *testing.T) {
	// Initialize
	d := map[string]interface{}{
		"test1": "message1",
		"test2": "message2",
	}
	h := handler{}

	// Encode
	c, e := h.serialize(d)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, "\x0e\xff\x81\x04\x01\x02\xff\x82\x00\x01\f\x01\x10\x00\x006\xff\x82\x00\x02\x05test1\x06string\f\n\x00\bmessage1\x05test2\x06string\f\n\x00\bmessage2", string(c))

	// Decode
	de := make(map[string]interface{})
	e = h.unserialize(c, &de)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, d, de)
}

// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cachemanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformKey(t *testing.T) {
	// Initialize
	h := handlerBase{
		prefix: "prefix_",
	}

	// Assert
	assert.Equal(t, "prefix_test", h.buildKey("test"))
}

func TestTransformTtl(t *testing.T) {
	// Initialize
	h := handlerBase{
		ttl: 5,
	}

	// Assert
	assert.Equal(t, int32(5), h.buildTTL(-1))
	assert.Equal(t, int32(3), h.buildTTL(3))
}

func TestSerialize(t *testing.T) {
	// Initialize
	d := map[string]interface{}{
		"test1": "message1",
		"test2": "message2",
	}
	h := handlerBase{}

	// Encode
	c, e := h.serialize(d)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, "{\"test1\":\"message1\",\"test2\":\"message2\"}\n", string(c))

	// Decode
	de := make(map[string]interface{})
	e = h.unserialize(c, &de)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, d, de)
}

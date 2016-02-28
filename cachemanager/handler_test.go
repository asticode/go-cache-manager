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

package cachemanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesTransformation(t *testing.T) {
	// Initialize
	d := []string{
		"test1",
		"test2",
	}

	// Encode
	c, e := ToBytes(d)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, "\f\xff\x81\x02\x01\x02\xff\x82\x00\x01\f\x00\x00\x10\xff\x82\x00\x02\x05test1\x05test2", string(c))

	// Decode
	de := []string{}
	e = FromBytes(c, &de)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, d, de)
}

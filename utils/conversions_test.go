package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	assert.Equal(t, 3, ParseInt("3"))
	assert.Equal(t, 42, ParseInt("42"))
	assert.Panics(t, func() { ParseInt("a") })
}

package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithListener(t *testing.T) {
	r := &Router{}

	testListener := ":1234"
	assert.Empty(t, r.listener)
	WithListener(testListener)(r)
	assert.Equal(t, testListener, r.listener)
}

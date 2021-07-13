package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultsAndOptionsAreUsed(t *testing.T) {
	r1, err := New()

	assert.NoError(t, err)
	assert.Equal(t, defaultListener, r1.listener)

	testListener := ":9999"
	assert.NotEqual(t, defaultListener, testListener)
	r2, err := New(WithListener(testListener))

	assert.NoError(t, err)
	assert.Equal(t, testListener, r2.listener)

}

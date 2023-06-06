package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask(t *testing.T) {
	s1, err := UnpackString("a4bc2d5e")
	assert.NoError(t, err)
	assert.Equal(t, s1, "aaaabccddddde")
	s2, err := UnpackString("abcd")
	assert.NoError(t, err)
	assert.Equal(t, s2, "abcd")
	s3, err := UnpackString("45")
	assert.Error(t, err)
	assert.Equal(t, s3, "")
}

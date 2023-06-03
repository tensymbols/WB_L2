package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	currTime, err := GetTime()
	assert.NoError(t, err)
	assert.Equal(t, currTime, time.Now().Format(time.TimeOnly))
}

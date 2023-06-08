package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestTask(t *testing.T) {

	data, err := exec.Command("go", "run", "./task.go", "-F", "input", "be").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, string(data), "be\n")

	data, err = exec.Command("go", "run", "./task.go", "-n", "-v", "input", "be").CombinedOutput()
	assert.NoError(t, err)
	fmt.Println(string(data))
	assert.Equal(t, string(data), "line 1:ek elsdf sdf\r\nline 2:fdsdfsfds dsfsdfsf\r\n")
}

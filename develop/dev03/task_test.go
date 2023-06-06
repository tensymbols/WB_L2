package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

type sortTest struct {
	Input    string
	Expected string
}

func TestTask(t *testing.T) {
	tests := []sortTest{
		{
			Input:    "kek b c d\nbek f g h\nz z\nl lll llll l",
			Expected: "bek f g h\nkek b c d\nl lll llll l\nz z",
		},
		{
			Input:    "a\ne\nc\nh\nb",
			Expected: "a\nb\nc\ne\nh",
		},
	}

	for i, v := range tests {
		fName := "in" + strconv.Itoa(i) + ".txt"
		f, err := os.Create(fName)
		assert.NoError(t, err)
		_, err = f.WriteString(v.Input)
		assert.NoError(t, err)
		err = f.Close()
		assert.NoError(t, err)
		exec.Command("go", "run", "./task.go", fName).Run()
		f, err = os.Open(fName)
		data, err := io.ReadAll(f)

		assert.NoError(t, err)
		assert.Equal(t, v.Expected, string(data))
		err = f.Close()
		assert.NoError(t, err)
	}
}

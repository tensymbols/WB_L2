package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	Input    []string
	Expected map[string][]string
}

func TestTask(t *testing.T) {
	tests := []testStruct{
		{
			[]string{
				"пятак", "пятка", "тяпка", "листок", "пятка", "слиток", "столик",
			},
			map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			[]string{
				"ток", "кот", "Кот", "лес", "Сел", "слиток", "столик",
			},
			map[string][]string{
				"слиток": {"слиток", "столик"},
				"кот":    {"кот", "ток"},
				"лес":    {"лес", "сел"},
			},
		},
	}
	t.Run("anagrams test", func(t *testing.T) {
		for _, v := range tests {
			anagrams := FindAnagrams(v.Input)
			assert.Equal(t, len(anagrams), len(v.Expected))
			for k := range anagrams {
				assert.Equal(t, anagrams[k], v.Expected[k])
			}
		}
	})

}

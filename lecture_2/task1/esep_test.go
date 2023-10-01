package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	test := []struct {
		name           string
		num            int
		expectedOutput string
	}{
		{
			name:           "test 1",
			num:            3,
			expectedOutput: "III",
		},
		{
			name:           "test 2",
			num:            58,
			expectedOutput: "LVIII",
		},
		{
			name:           "test 3",
			num:            1994,
			expectedOutput: "MCMXCIV",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := intToRoman(tt.num)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

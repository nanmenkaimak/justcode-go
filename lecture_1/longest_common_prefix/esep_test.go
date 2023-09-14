package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testEx := []struct {
		name           string
		strs           []string
		expectedOutput string
	}{
		{
			name:           "test 1",
			strs:           []string{"flower", "flow", "flight"},
			expectedOutput: "fl",
		},
		{
			name:           "test 2",
			strs:           []string{"dog", "racecar", "car"},
			expectedOutput: "",
		},
	}

	for _, tt := range testEx {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.strs)

			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

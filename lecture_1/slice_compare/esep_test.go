package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceCompare(t *testing.T) {
	testEx := []struct {
		name           string
		arr1           []int
		arr2           []int
		expectedOutput bool
	}{
		{
			name:           "correct",
			arr1:           []int{1, 2, 3},
			arr2:           []int{2, 1, 3},
			expectedOutput: true,
		},
		{
			name:           "different length",
			arr1:           []int{1, 2, 3, 5},
			arr2:           []int{2, 1, 3},
			expectedOutput: false,
		},
		{
			name:           "different values",
			arr1:           []int{1, 2, 3},
			arr2:           []int{2, 1, 5},
			expectedOutput: false,
		},
	}

	for _, tt := range testEx {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceCompare(tt.arr1, tt.arr2)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testEx := []struct {
		name           string
		nums           []int
		target         int
		expectedOutput []int
	}{
		{
			name:           "test 1",
			nums:           []int{2, 7, 11, 15},
			target:         9,
			expectedOutput: []int{0, 1},
		}, {
			name:           "test 2",
			nums:           []int{3, 2, 4},
			target:         6,
			expectedOutput: []int{1, 2},
		}, {
			name:           "test 3",
			nums:           []int{3, 3},
			target:         6,
			expectedOutput: []int{0, 1},
		},
	}

	for _, tt := range testEx {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)

			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

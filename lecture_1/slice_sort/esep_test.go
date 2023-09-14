package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSorting(t *testing.T) {
	testEx := []struct {
		name           string
		sortingStruct  sliceSorting
		expectedOutput []int
	}{
		{
			name: "test 1",
			sortingStruct: sliceSorting{
				values: []int{4, 5, 3, 6, 2, 1},
			},
			expectedOutput: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "test 2",
			sortingStruct: sliceSorting{
				values: []int{1, 2, 3, 4, 5, 6},
			},
			expectedOutput: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range testEx {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.sortingStruct.sorting()

			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

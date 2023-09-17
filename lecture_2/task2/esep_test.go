package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	test := []struct {
		name           string
		list1          []int
		list2          []int
		expectedOutput []int
	}{
		{
			name:           "test 1",
			list1:          []int{1, 2, 4},
			list2:          []int{1, 3, 4},
			expectedOutput: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:           "test 2",
			list1:          []int{},
			list2:          []int{},
			expectedOutput: nil,
		},
		{
			name:           "test 3",
			list1:          []int{},
			list2:          []int{0},
			expectedOutput: []int{0},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := listToSlice(mergeTwoLists(sliceToList(tt.list1), sliceToList(tt.list2)))

			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

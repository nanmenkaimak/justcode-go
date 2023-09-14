package main

import "sort"

func sliceCompare(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

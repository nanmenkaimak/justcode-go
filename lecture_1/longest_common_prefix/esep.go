package main

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	n := len(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[n-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
}

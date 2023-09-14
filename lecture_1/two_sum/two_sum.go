package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i, num := range nums {
		m[num] = i
	}

	for indX, num := range nums {
		if indY, ok := m[target-num]; ok && indX != indY {
			return []int{indX, indY}
		}
	}

	return []int{}
}

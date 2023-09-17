package main

func intToRoman(num int) string {
	ans := ""
	b := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	a := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	i := 0
	for num > 0 {
		for a[i] <= num {
			num -= a[i]
			ans += b[i]
		}
		i++
	}
	return ans
}

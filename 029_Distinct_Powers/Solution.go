package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(distinctPowerTerms(n))
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func distinctPowerTerms(n int) int {
	count := 0
	test := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if test[i] {
			continue
		}
		combined := make(map[int]struct{})
		pw := 2
		for power(i, pw) <= n {
			test[power(i, pw)] = true
			for b := 2; b <= n; b++ {
				if b*pw > n {
					combined[b*pw] = struct{}{}
				}
			}
			pw += 1
		}
		count += len(combined) + n - 1
	}
	return count
}
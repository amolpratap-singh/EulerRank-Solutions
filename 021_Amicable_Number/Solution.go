package main

import (
	"fmt"
)

const MAX_NUM = 100000

func main() {
	numDivisors := make([]int, MAX_NUM+1)
	divisorSums := make([]int, MAX_NUM+1)

	for i := 0; i < len(numDivisors); i++ {
		numDivisors[i] = 1
	}
	numDivisors[0], numDivisors[1] = 0, 0

	for factor := 2; factor <= int(MAX_NUM/2); factor++ {
		for num := factor * 2; num <= MAX_NUM; num += factor {
			numDivisors[num] += factor
		}
	}

	for i := 1; i <= MAX_NUM; i++ {
		divisorSums[i] = divisorSums[i-1]
		if numDivisors[i] <= MAX_NUM && numDivisors[i] != i &&
			numDivisors[numDivisors[i]] == i {
			divisorSums[i] += i
		}
	}

	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(divisorSums[n])
	}
}

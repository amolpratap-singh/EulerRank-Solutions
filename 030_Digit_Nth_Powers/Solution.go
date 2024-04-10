package main

import (
	"fmt"
)

func main() {
	var nth_pow int
	fmt.Scan(&nth_pow)

	output := make(map[int]bool)

	for i := 100; i < 1000000; i++ {
		_sum := 0
		num := i
		for num != 0 {
			digit := num % 10
			_sum += pow(digit, nth_pow)
			num /= 10
		}
		if _sum == i {
			output[i] = true
		}
	}

	result := 0
	for num := range output {
		result += num
	}
	fmt.Println(result)
}

func pow(x, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

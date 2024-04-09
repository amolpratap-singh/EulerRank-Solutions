package main

import (
	"fmt"
	"math"
)

func pentagonNum(num int) int {
	return num * (3*num - 1) / 2
}

func isPentagonNum(numP int) bool {
	// Using Quadratic eqn
	a := 3
	b := -1
	c := -2
	D := int(math.Sqrt(float64(b*b - 4*a*c*numP)))
	x := 2 * a
	rt1 := (-b + D) / x
	rt2 := (-b - D) / x

	if rt1 > 0 && numP == pentagonNum(rt1) {
		return true
	} else if rt2 > 0 && numP == pentagonNum(rt2) {
		return true
	} else {
		return false
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	for num := k + 1; num < n; num++ {
		if isPentagonNum(pentagonNum(num) - pentagonNum(num - k)) || isPentagonNum(pentagonNum(num) + pentagonNum(num - k)) {
			fmt.Println(pentagonNum(num))
		}
	}
}

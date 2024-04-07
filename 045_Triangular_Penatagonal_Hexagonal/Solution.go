package main

import (
	"fmt"
	"math"
)

func isTriangular(x int) bool {
	return math.Sqrt(float64(1+8*x)) == math.Floor(math.Sqrt(float64(1+8*x)))
}

func isPentagonal(x int) bool {
	return math.Sqrt(float64(1+24*x)) == math.Floor(math.Sqrt(float64(1+24*x))) && int(math.Sqrt(float64(1+24*x))+1)%6 == 0
}

func main() {
	var n, k, b int
	fmt.Scan(&n, &k, &b)

	var penMax, hexMax int
	if k == 3 {
		penMax = int((math.Sqrt(float64(1+24*n)) + 1) / 6)
		for i := 1; i < penMax; i++ {
			p := i * (3*i - 1) / 2
			if isTriangular(p) {
				fmt.Println(p)
			}
		}
	} else {
		hexMax = int((math.Sqrt(float64(1+8*n)) + 1) / 4)
		for i := 1; i < hexMax; i++ {
			h := i * (2*i - 1)
			if isPentagonal(h) {
				fmt.Println(h)
			}
		}
	}
}

package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func combinations(n, r int) *big.Int {
	numerator := factorial(n)
	denominator := new(big.Int).Mul(factorial(r), factorial(n-r))
	return new(big.Int).Div(numerator, denominator)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	threshold := big.NewInt(int64(k))
	count := 0
	for i := 1; i <= n; i++ {
		for r := 0; r < n; r++ {
			if combinations(i,r).Cmp(threshold) > 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}

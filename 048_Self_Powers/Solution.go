package main

import (
	"fmt"
	"math/big"
)

func main() {
	var num int
	fmt.Scan(&num)

	totalSum := new(big.Int)
	modulus := new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil)

	for i := 1; i <= num ; i++ {
		iPow := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil)
		totalSum.Add(totalSum, iPow)
	}

	lastTenDigits := new(big.Int).Mod(totalSum, modulus)
	fmt.Println(lastTenDigits)

}
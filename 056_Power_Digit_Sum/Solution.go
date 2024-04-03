package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)

	maxSum := 0

	for i := 1; i<num; i++ {
		for j := 1; j<num; j++ {
			if maxSum < getPowerSumOfDigits(i,j) {
				maxSum = getPowerSumOfDigits(i,j)
			}
		}
	}

	fmt.Println(maxSum)

}

func getPowerSumOfDigits(a, b int) int {
	pow := new(big.Int).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
	sum := 0
	for _, char := range pow.String() {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting character to integer:", err)
			continue
		}
		sum += digit
	}

	return sum
}
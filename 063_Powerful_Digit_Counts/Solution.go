package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	i := 1
	for {
		num := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(int64(n)), nil)
		if len(strconv.Itoa(int(num.Int64()))) == n {
			fmt.Println(num)
		}
		if len(strconv.Itoa(int(num.Int64()))) > n {
			break
		}
		i++
	}
}

package main

import (
	"fmt"
)

func main() {
	var testCase int
	fmt.Scan(&testCase)

	for i := 0; i < testCase; i++ {
		var n int
		fmt.Scan(&n)
		n = (n - 1) / 2
        result := (16 * n * n * n + 30 * n * n + 26 * n + 3) / 3 % (int(1e9) + 7)
        fmt.Println(result)

	}
}
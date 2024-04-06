package main

import (
	"fmt"
	"strconv"
)

func convertToBase(decimalNumber, base int) string {
	if decimalNumber == 0 {
		return "0"
	}

	digits := "0123456789"
	var result string
	for decimalNumber > 0 {
		remainder := decimalNumber % base
		result = string(digits[remainder]) + result
		decimalNumber /= base
	}
	return result
}

func checkPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	for i, j := 0, len(numStr)-1; i < j; i, j = i+1, j-1 {
		if numStr[i] != numStr[j] {
			return false
		}
	}
	return true
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, k int
	fmt.Scan(&n, &k)

	var sum int
	for num := 1; num < n; num++ {
		baseRepr := convertToBase(num, k)
		baseReprInt, err := strconv.Atoi(baseRepr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if checkPalindrome(num) && checkPalindrome(baseReprInt) {
			sum += num
		}
	}
	fmt.Println(sum)
}

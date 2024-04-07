package main

import (
	"fmt"
	"strings"
)

var numbers = map[int]string{
	0: "", 1: " One", 2: " Two", 3: " Three", 4: " Four", 5: " Five",
	6: " Six", 7: " Seven", 8: " Eight", 9: " Nine",
}

var teens = map[int]string{
	10: " Ten", 11: " Eleven", 12: " Twelve", 13: " Thirteen",
	14: " Fourteen", 15: " Fifteen", 16: " Sixteen",
	17: " Seventeen", 18: " Eighteen", 19: " Nineteen",
}

var tens = map[int]string{
	20: " Twenty", 30: " Thirty", 40: " Forty", 50: " Fifty",
	60: " Sixty", 70: " Seventy", 80: " Eighty", 90: " Ninety",
}

func convertNumWords(n int) string {
	if n < 10 {
		return numbers[n]
	} else if n < 20 {
		return teens[n]
	} else if n < 100 && n%10 == 0 {
		return tens[n]
	} else if n < 100 {
		return tens[10*(n/10)] + numbers[n%10]
	} else if n < 1000 {
		return numbers[n/100] + " Hundred" + convertNumWords(n%100)
	} else if n < 1000000 {
		return convertNumWords(n/1000) + " Thousand" + convertNumWords(n%1000)
	} else if n < 1000000000 {
		return convertNumWords(n/1000000) + " Million" + convertNumWords(n%1000000)
	} else {
		return convertNumWords(n/1000000000) + " Billion" + convertNumWords(n%1000000000)
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var num int
		fmt.Scan(&num)
		output := convertNumWords(num)
		fmt.Println(strings.TrimSpace(output))
	}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var testCase int
	fmt.Scan(&testCase)

	namesList := make([]string, testCase)
	for i := 0; i < testCase; i++ {
		var name string
		fmt.Scan(&name)
		namesList[i] = name
	}

	sort.Strings(namesList)

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var nameQ string
		fmt.Scan(&nameQ)

		sum := 0
		for count, name := range namesList {
			if nameQ == name {
				for _, k := range name {
					sum += int(k - 'A' + 1)
				}
				sum *= count + 1
			}
		}
		fmt.Println(sum)
	}
}

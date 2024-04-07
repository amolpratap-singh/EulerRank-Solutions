package main

import (
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	longestNum := []int32{0, 1}
	longest := int32(0)
	terms := []int32{0, 1}

	var t, n int
	fmt.Scan(&t)
	for tItr := 0; tItr < int(t); tItr++ {
		fmt.Scan(&n)
		for nItr := len(longestNum); nItr <= n; nItr ++ {
			num := nItr
			count := int32(0)
			for (num != 1 && num >= nItr) {
				if num % 2 == 0 {
					num /= 2
				} else {
					num = 3 * num + 1
				}
				count += 1
			}

			length := count + terms[num]
			terms = append(terms, length)
			if length >= longest {
				longest = length
				longestNum = append(longestNum, int32(nItr))
			} else {
				longestNum = append(longestNum, longestNum[len(longestNum) -1])
			}
		}
		fmt.Println(longestNum[n])
	}
	
}
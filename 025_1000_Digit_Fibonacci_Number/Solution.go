package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	digits := []int{0, 1}
	f1, f2 := 1, 1
	c := 10
	for i := 3; i < 25000; i++ {
		f1, f2 = f2, f1+f2
		if f2 >= c {
			c *= 10
			digits = append(digits, i)
		}
	}

	var testCase int
	fmt.Scan(&testCase)

	for i := 0; i < testCase; i++ {
		var num int
		fmt.Scan(&num)
		fmt.Println(digits[num])

	}

}

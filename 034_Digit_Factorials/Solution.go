package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	sumNumber := 0
	for tItr := 10; tItr <= int(t); tItr++ {
		sum := 0
		num := tItr

		for num > 0 {
			sum += factorialRecursive(num % 10)
			num = int(num / 10)
		}

		if (sum % tItr == 0) {
			sumNumber += tItr
		}	
	}
	fmt.Println(sumNumber)
}

func factorialRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

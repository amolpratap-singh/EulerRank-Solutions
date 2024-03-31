package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
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

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int64(nTemp)
		factNum := factorialNum(int(n))
		sum := 0

		for _, digitChar := range factNum.String() {
			digit, _ := strconv.Atoi(string(digitChar))
			sum += digit
		}

		fmt.Println(sum)
	}
}

func factorialNum(n int) *big.Int {
	res := big.NewInt(1)

	for i := 2; i <= n; i++ {
		res.Mul(res, big.NewInt(int64(i)))
	}

	return res
}

/*
// recursion function works only till 10 not more than above due to int64 limit
func factorialNum(num int64) int64 {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorialNum(num-1)
}
*/

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

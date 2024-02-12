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
	sum := new(big.Int)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp := strings.TrimSpace(readLine(reader))
		num := new(big.Int)
		val, check := num.SetString(nTemp, 10)

		if !check {
			fmt.Println("Failed to parse Integer")
			return
		}

		sum.Add(sum, val)

	}
	fmt.Println(sum.String()[:10])
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

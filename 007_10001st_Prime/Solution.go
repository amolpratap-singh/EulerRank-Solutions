// Only 3 cases get pass out of 5 due to Runtime error for two cases

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const N = 10000

func sieveOfEratosthenes() []int32 {
	primeNum := make([]bool, N+1)
	for i := int32(2); i*i <= N; i++ {
		if primeNum[i] == false {
			for j := i * i; j <= N; j += i {
				primeNum[j] = true
			}
		}
	}
	prime := make([]int32, 0)
	for i := int32(2); i <= N; i++ {
		if primeNum[i] == false {
			prime = append(prime, i)
		}
	}
	return prime
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		fmt.Println(sieveOfEratosthenes()[n-1])

	}
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

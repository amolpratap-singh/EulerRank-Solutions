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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		var largestPalindrome int32

		for largestPalindrome == 0 {
			n--
			if checkPalindrome(n) {
				for j := int32(100); j < 1000; j++ {
					r := float64(n) / float64(j)
					if r == float64(int(r)) && r < 1000 {
						largestPalindrome = n
						break
					}
				}
			}
		}

		fmt.Println(largestPalindrome)

	}
}

func checkPalindrome(num int32) bool {
	strNum := strconv.Itoa(int(num))
	for i := 0; i < len(strNum)/2; i++ {
		if strNum[i] != strNum[len(strNum)-1-i] {
			return false
		}
	}
	return true
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

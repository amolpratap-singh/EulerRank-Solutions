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
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)
		fmt.Println(getMaxProduct(int(n)))
    }
}

func getMaxProduct(num int) int {
	product := -1
	var b, c int
	for a := 1; a <= num/3; a++ {
		b = (num*num - 2*num*a) / (2*num - 2*a)
		c = num - b - a

		if c*c == (a*a + b*b) {
			temp := a * b * c
			if temp > product {
				product = temp
			}
		}
	}
	return product
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

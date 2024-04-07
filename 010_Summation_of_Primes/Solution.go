package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"math"
	"sort"
)

const LIMIT = 2000000


func main() {

	sums := primeSums()
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)


    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)
		index := sort.Search(len(sums), func(idx int) bool {
			return int32(sums[idx]) > n
		})
		fmt.Println(sums[index])

    }
}

func primeSums() []int {
	slice := make([]bool, LIMIT)
	primes := []int{}
	sums := []int{0}

	for num := 2; num < int(math.Sqrt(float64(LIMIT))); num++ {
		if !slice[num] {
			primes = append(primes, num)
			sums = append(sums, sums[len(sums)-1]+num)
			for mul := num * num; mul < LIMIT; mul += num {
				slice[mul] = true
			}
		}
	}

	for number := int(math.Sqrt(float64(LIMIT))); number < LIMIT; number++ {
		if !slice[number] {
			primes = append(primes, number)
			sums = append(sums, sums[len(sums)-1]+number)
		}
	}

	return sums

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

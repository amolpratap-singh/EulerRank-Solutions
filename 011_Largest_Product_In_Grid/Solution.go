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

	var grid [][]int32
	for i := 0; i < 20; i++ {
		gridRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var gridRow []int32
		for _, gridRowItem := range gridRowTemp {
			gridItemTemp, err := strconv.ParseInt(gridRowItem, 10, 64)
			checkError(err)
			gridItem := int32(gridItemTemp)
			gridRow = append(gridRow, gridItem)
		}

		if len(gridRow) != 20 {
			panic("Bad input")
		}

		grid = append(grid, gridRow)
	}

	var maxProduct int32 = 0
	var product int32 = 0

	for row := 0; row < 20; row++ {
		for col := 0; col < 20; col++ {
			// horizontal product
			if col+3 < 20 {
				product = grid[row][col] * grid[row][col+1] * grid[row][col+2] * grid[row][col+3]
				if product > maxProduct {
					maxProduct = product
				}
			}

			// vertical product
			if row+3 < 20 {
				product = grid[row][col] * grid[row+1][col] * grid[row+2][col] * grid[row+3][col]
				if product > maxProduct {
					maxProduct = product
				}
			}

			// Diagonal product
			if col+3 < 20 && row+3 < 20 {
				product = grid[row][col] * grid[row+1][col+1] * grid[row+2][col+2] * grid[row+3][col+3]
				product = grid[row][col+3] * grid[row+1][col+2] * grid[row+2][col+1] * grid[row+3][col]

				if product > maxProduct {
					maxProduct = product
				}
			}
		}
	}

	fmt.Println(maxProduct)

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

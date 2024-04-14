package main

import (
	"fmt"
	"math"
)

func area(x1, y1, x2, y2, x3, y3 int) float64 {
	return math.Abs(float64((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))) / 2.0)
}

func main() {
	var n, x1, y1, x2, y2, x3, y3 int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

		triangle1Area := area(0, 0, x2, y2, x3, y3)
		triangle2Area := area(x1, y1, 0, 0, x3, y3)
		triangle3Area := area(x1, y1, x2, y2, 0, 0)

		triangleArea := area(x1, y1, x2, y2, x3, y3)

		if triangleArea == triangle1Area+triangle2Area+triangle3Area {
			count++
		}
	}

	fmt.Println(count)
}

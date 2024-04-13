package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var expoList []struct{ log, base, exp float64 }
	var n, k int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var b, e int
		fmt.Scan(&b, &e)
		logVal := float64(e) * math.Log(float64(b))
		expoList = append(expoList, struct{ log, base, exp float64 }{logVal, float64(b), float64(e)})
	}

	fmt.Scan(&k)

	sort.Slice(expoList, func(i, j int) bool {
		return expoList[i].log < expoList[j].log
	})

	x := expoList[k-1]
	fmt.Println(int(x.base), int(x.exp))
}

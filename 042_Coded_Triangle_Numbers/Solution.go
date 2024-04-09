package main

import (
    "fmt"
    "math"
)

func main() {
    var t int
    fmt.Scan(&t)

    for i := 0; i < t; i++ {
        var N int
        fmt.Scan(&N)
        
        a := 1
        b := 1
        c := -2
        
        D := int(math.Sqrt(float64(b*b - 4*a*c*N)))
        
        rt1 := (-b + D) / (2 * a)
        rt2 := (-b - D) / (2 * a)
        
        if rt1 > 0 && N == (rt1 * (rt1 + 1) / 2) {
            fmt.Println(rt1)
        } else if rt2 > 0 && N == (rt2 * (rt2 + 1) / 2) {
            fmt.Println(rt2)
        } else {
            fmt.Println(-1)
        }
    }
}

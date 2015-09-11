package main

import (
    "fmt"
    "math"
)

type MyFloatType float64

func (f MyFloatType) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    f := MyFloatType(-math.Sqrt2)
    fmt.Println(f.Abs())
}

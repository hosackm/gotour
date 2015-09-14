package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

type Vertex struct {
    X, Y float64
}

type MyFloat float64

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f
    a = &v
    //a = v // Not a *Vertex, so won't work

    fmt.Println(a.Abs())
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (m MyFloat) Abs() float64 {
    if m < 0 {
        return float64(-m)
    }
    return float64(m)
}

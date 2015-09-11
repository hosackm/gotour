package main

import (
    "fmt"
)

func Pic(dx, dy int) (p [][]uint8) {
    p = make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        p[i] = make([]uint8, dx)
        for j := 0; j < dx; j++ {
            p[i][j] = uint8(i^j)
        }
    }
    return /* naked return of p */
}


func main() {
    fmt.Println(Pic(8, 10))
}
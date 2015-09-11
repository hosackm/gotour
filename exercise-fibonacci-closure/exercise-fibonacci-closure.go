package main

import (
    "fmt"
    "strings"
)

func fibonacci() func() int {
    x, y := 0, 1
    return func() int {
        x, y = y, x+y
        return x
    }
}

func main() {
    f := fibonacci()
    num := 20
    s := "The First %d Fibonacci Numbers\n"
    fmt.Printf(s, num)
    fmt.Println(strings.Repeat("=", len(s)))
    for i := 0; i < num; i++ {
        fmt.Println(f())
    }
}

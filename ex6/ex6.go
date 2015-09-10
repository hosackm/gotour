package main

import (
    "fmt"
)

/*
Essentially the same function prototype but longer:
func add(x int, y int) int {
*/
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(3, 5))
}

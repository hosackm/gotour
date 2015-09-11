package main

import "fmt"

func main() {
    a := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(a)

    for i := 0; i < len(a); i++ {
        fmt.Printf("a[%d] == %d\n", i, a[i])
    }
}

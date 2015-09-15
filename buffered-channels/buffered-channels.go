package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    //Uncomment this for buffer overflow
    //ch <- 3
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

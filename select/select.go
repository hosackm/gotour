package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select{
            case c <- x:
                x, y = y, x+y
            case <-quit:
                fmt.Println("quit")
                return
        }
    }
}

func main() {
    c := make(chan int, 10)
    quit := make(chan int)

    //Thread 1
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
    quit <- 0
    }()

    //Thread 2
    fibonacci(c, quit)
}

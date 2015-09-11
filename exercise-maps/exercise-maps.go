package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "Hello this is a sentence to count"
    fmt.Printf("Counting '%s'\n", s)
    fmt.Println(WordCount(s))
}

func WordCount(s string) (count map[string]int) {
    count = make(map[string]int)
    words := strings.Fields(s)
    for _, word := range words {
        count[word]++
    }    

    return
}

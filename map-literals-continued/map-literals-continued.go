package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    /* Don't have to specify Vertex{Lat, Long} inside literal */
    "Bell Labs": {40.68433, -74.1239},
    "Google": {37.42202, -122.08408},
}

func main() {
    fmt.Println(m)
}

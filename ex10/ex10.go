package main

import (
    "fmt"
)

/* Initializers with type specified */
var i, j int = 1, 2

func main() {
    /* If type not specified, variable infers from initializer */
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}

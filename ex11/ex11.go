package main

import (
    "fmt"
)

/* Cannot use ':=' outside of a function
    x := 1234 // DONT DO THIS!
 */

func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}

package main

import "fmt"

type IPAddr [4]byte

func main() {
    addrs := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for n, a := range addrs {
        fmt.Printf("%v: %v\n", n, a)
    }
}

func (a IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
}

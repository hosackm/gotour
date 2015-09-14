package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

func (r *rot13Reader) Read(b []byte) (int, error) {
    n, err := r.r.Read(b)
    var c byte
    for i := 0; i < len(b); i++ {
        c = b[i]
        if((c > 'A' && c < 'N') || (c > 'a' && c < 'n')) {
            b[i] = c + 13
        } else if((c > 'A' && c < 'N') || (c > 'a' && c < 'n')) {
            b[i] = c - 13
        }
    }
    return n, err
}

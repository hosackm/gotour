package main

import(
    "fmt"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func _walk(t *tree.Tree, ch chan int) {
    if t != nil{
        _walk(t.Left, ch)
        ch <- t.Value
        _walk(t.Right, ch)
    }
}

// Helper function to close channel after walk finishes
func Walk(t *tree.Tree, ch chan int) {
    _walk(t, ch)
    close(ch)   
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    
    for i := range ch1 {
        if i != <-ch2 {
            return false
        }
    }
    
    return true
}

func main() {
    ch := make(chan int, 10)
    go Walk(tree.New(1), ch)
    
    for i := range ch {
        fmt.Println(i)
    }   
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}

package main

import "fmt"

func Generate() chan int {
    ch := make(chan int)
    go func() {
        for i := 2; ; i++ {
            ch <- i
        }
    }()
    return ch
}

func Filter(in <- chan int, prim int) chan int{
    out := make(chan int)
    go func() {
        for {
            if i := <- in; i%prim != 0 {
                out <- i
            }
        }
    }()
    return out
}

func main(){
    in := Generate()
    for i := 0; i < 100; i++ {
        prim := <- in
        fmt.Printf("[%d] --> %d\n", i, prim)
        in = Filter(in, prim)
    }
}


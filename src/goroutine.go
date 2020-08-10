package main

import (
    "fmt"
)


func fibonacci(n int, ch chan int){
    x, y := 1,1
    for i := 0; i < n; i++{
        ch <- x
        x, y = y, x + y
    }
    close(ch)
}


func main(){
    c := make(chan int, 10)
    fibonacci(cap(c), c)
    for value := range c{
        fmt.Printf("%d\n", value)
    }
}

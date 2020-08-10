package main

import "fmt"

func sum_and_product(a, b int)(int, int){
    return a+b, a*b
}

func main(){
    x := 3
    y := 4
    
    sum, product := sum_and_product(x, y)

    fmt.Printf("%d + %d = %d\n", x, y, sum)
    fmt.Printf("%d * %d = %d\n", x, y, product)
}

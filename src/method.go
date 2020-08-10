package main
import (
    "math"
    "fmt"
)

type Rectangle struct{
    width, height float64
}

type Circle struct{
    radius float64
}

func (rect Rectangle) area() float64{
    return rect.width * rect.height
}

func (cir Circle) area() float64{
    return cir.radius * cir.radius * math.Pi;
}

func main(){
    rect := Rectangle{width:5, height:10}
    circle := Circle{radius:3}

    fmt.Println("area of rect:", rect.area())
    fmt.Println("area of circle:", circle.area())
}

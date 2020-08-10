package main

import (
    "fmt"
    "strconv"
)

type Person struct{
    name string
    age  int
    phone string
}


func (h Person) String()string{
    return "Name:" + h.name + " Age:" + strconv.Itoa(h.age) + " Phone:" + h.phone
}

type Element interface{}
type List []Element

func main(){
    list := make(List, 3)
    list[0] = 1
    list[1] = "hello world"
    list[2] = Person{name:"bob", age:18, phone:"123456"}

    for index, element := range(list){
        switch value := element.(type){
            case int:
                fmt.Printf("list[%d] is an int and its value is : %d\n", index, value)
            case string:
                fmt.Printf("list[%d] is a string and its value is : %s\n", index, value)
            case Person:
                fmt.Printf("list[%d] is a Person and its value is : %s\n", index, value)
            default:
                fmt.Println("list[%d] is of a different type\n", index)
        }
    }
}

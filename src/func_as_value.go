package  main
import "fmt"


type testInt func(int) bool

func is_odd(n int) bool{
    return 1 == n % 2
}

func is_even(n int) bool{
    return 0 == n % 2
}

func filter(slice []int, f testInt) [] int{
    var result []int
    for _, value := range slice {
        if f(value){
            result = append(result, value)
        }
    }

    return result
}


func main(){
    slice := []int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println("slice: ", slice);

    odd := filter(slice, is_odd)
    even := filter(slice, is_even)
    fmt.Println("odd : ", odd)
    fmt.Println("even : ", even)
}

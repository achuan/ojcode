package main

import(
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func isPowerOfFour(num int) bool {
    if (num & (num-1) == 0) && (num & 0xAAAAAAAA) != num {
        return true
    }

    return false;
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    for true {
        str, _, _ := reader.ReadLine()
        num, _  := strconv.Atoi(string(str))
        fmt.Println(string(str), isPowerOfFour(num))
    }
}

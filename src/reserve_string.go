package main

import(
    "fmt"
    "os"
    "bufio"
)

func reverseString(s string) string {
    str := []byte(s)
    n := len(str)
    for i:= 0; i < (n+1)/2; i++{
        str[i], str[n-i-1] = str[n-i-1], str[i]
    }

    return string(str)
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    for true {
        str,_, _ := reader.ReadLine()
        fmt.Println(string(str), reverseString(string(str)));
    }
}

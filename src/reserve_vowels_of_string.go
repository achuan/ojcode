/*
 *https://leetcode.com/problems/reverse-vowels-of-a-string/
 *Description:
 *Write a function that takes a string as input and reverse only the vowels of a string.
 *Example 1:
 *Given s = "hello", return "holle".
 *Example 2:
 *Given s = "leetcode", return "leotcede".
 */

package main

import (
    "fmt"
    "bufio"
    "os"
)

func vowels(ch byte) bool{
    if  ch == 'a' || ch == 'o' || ch == 'e' || ch == 'i' || ch == 'u' ||ch == 'A' || ch == 'O' || ch == 'E' || ch == 'I' || ch == 'U' {
        return true;
    }

    return false;
}

func reverseVowels(str string) string {
    s := []byte(str)
    len := len(s)
    r := len - 1
    for l :=0 ; l<r; l++{
        if vowels(s[l]){
            for r > l && !vowels(s[r]) {
                r--
            }

            if l == r {
                break;
            }

            s[l], s[r] = s[r], s[l]
            r--
        }
    }

    return  string(s)
}

func main(){
    running := true
    reader := bufio.NewReader(os.Stdin)

    for running {
        str, _, _ := reader.ReadLine()   
        fmt.Println("string is ", string(str), reverseVowels(string(str)));
        if string(str) == "quit" {
            running = false;
        }
    }
}

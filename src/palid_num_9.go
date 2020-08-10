//https://leetcode.com/problems/palindrome-number/description/
package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	tmp := x
	sum := 0
	for x > 0 {
		sum = sum*10 + x%10
		x /= 10
	}
	if tmp == sum {
		return true
	}
	return false
}

func main() {
	array_int := []int{0, 1, 10, 15, 20, 111, 100, 123, 10000, 100001, 12321, 123321}
	for _, item := range array_int {
		fmt.Println(item, "is param:", isPalindrome(item))
	}
}

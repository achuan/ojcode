//https://leetcode.com/problems/valid-perfect-square/description/
package main

import "fmt"

func isPerfectSquare(num int) bool {
	left := 0
	right := num

	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	for i := 0; i < 200; i++ {
		fmt.Println(i, " is perfect square: ", isPerfectSquare(i))
	}
}

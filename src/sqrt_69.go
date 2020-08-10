//https://leetcode.com/problems/sqrtx/description/
package main

import "fmt"

func mySqrt(x int) int {
	left := 0
	right := x

	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

func main() {
	for i := 0; i < 200; i++ {
		fmt.Println(i, " sqrt = ", mySqrt(i))
	}
}

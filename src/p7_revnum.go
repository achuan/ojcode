package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/reverse-integer/
const MaxInt = 1<<31 - 1
const MinInt = -1 << 31

func reverse(x int) int {
	var num int
	for x != 0 {
		pop := x % 10

		if num > MaxInt/10 || (num == MaxInt/10 && pop > 7) {
			return 0
		}
		if num < MinInt/10 || (num == MinInt/10 && pop < -8) {
			return 0
		}

		num = num*10 + pop
		x = x / 10
	}

	return num
}

func main() {
	testArray := []int{-123, 122, 121, 343, -1, -145, 1234, 9080, 0}

	for _, item := range testArray {
		fmt.Printf("reverse of %v is %v\n", item, reverse(item))
	}
}

package main

import "fmt"

const INT_MAX = 1 << 31 - 1

//https://leetcode-cn.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	n := 0
	sign := 1
	if divisor < 0 {
		divisor = -divisor
		sign = -sign
	}
	if dividend < 0 {
		dividend = -dividend
		sign = -sign
	}
	for divisor<<n <= dividend {
		if divisor << n == dividend{
			if sign > 0 && n == 31 {
				return INT_MAX
			}
			return sign * (1 << n)
		}
		n++
	}
	if n == 0 {
		return 0
	}

	left, right := 1 << (n-1), 1 << n
	for left <= right {
		mid := (left + right) / 2
 		target := mid * divisor
		if target == dividend {
			return sign * mid
		} else if target < dividend {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return sign * (left - 1)
}

func main() {
	testArray := []struct {
		dividend int
		divisor int
	}{
		{-2147483648, 1},{10,3},{20,2},{19,3},{18,3},{17,3},{16,3}, {15,3},
	}
	for _, item := range testArray{
		fmt.Printf("result of %v / %v = %v\n", item.dividend, item.divisor, divide(item.dividend, item.divisor))
	}
}

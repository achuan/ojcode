package main

import "fmt"

//https://leetcode-cn.com/problems/string-to-integer-atoi/
const MaxInt = 1<<31 - 1
const MinInt = -1 << 31

func myAtoi(s string) int {
	target := 0
	sign := 0
	for _, ch := range s {
		if sign == 0 {
			if ch == ' ' {
				continue
			} else if ch == '-' {
				sign = -1
				continue
			} else if ch == '+' {
				sign = 1
				continue
			}
		}
		if sign == 0 {
			sign = 1
		}
		if ch < '0' || ch > '9' {
			break
		}
		pop := int(ch - '0')
		if target > MaxInt/10 || (target == MaxInt/10 && pop > 7) {
			return MaxInt
		}
		if target < MinInt/10 || (target == MinInt/10 && pop > 8) {
			return MinInt
		}

		target = target*10 + pop*sign
	}
	return target
}

func main() {
	testArray := []string{"9999999999999", "122", "+1234", "  123",
		"  +123", "  -123", " -123", "-999999999999",
		"-2147483648", "-2147483649",
		"2147483647", "2147483648",
		"2147483646", "-2147483647", "+-99", "-+99", "words and 987"}

	for _, item := range testArray {
		fmt.Printf("reverse of %v is %v\n", item, myAtoi(item))
	}
}

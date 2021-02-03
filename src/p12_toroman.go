package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/integer-to-roman/

func intToRoman(num int) string {
	var romanStr string
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for index, value := range values {
		if num/value > 0 {
			romanStr += strings.Repeat(romans[index], num/value)
			num %= value
		}
	}

	return romanStr
}

func main() {
	testArray := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 15, 19, 20, 39, 40, 50, 89, 90, 100, 500, 900, 1000, 3819}
	for _, item := range testArray {
		fmt.Printf("test result %v is %v\n", item, intToRoman(item))
	}

}

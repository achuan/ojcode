//https://leetcode.com/problems/add-digits/description/
package main

import "fmt"

func addDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	if sum < 10 {
		return sum
	}

	return addDigits(sum)
}

//https://en.wikipedia.org/wiki/Digital_root#Congruence_formula
func addDigits(num int) int {
	return 1 + (num-1)%9
}

func main() {
	num := 9999999999
	fmt.Printf("num:%v, digit:%v\n", num, addDigits(num))
}

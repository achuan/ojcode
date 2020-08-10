//https://leetcode.com/problems/ugly-number/description/
package main

import "fmt"

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	base := []int{2, 3, 5}
	for _, prime := range base {
		for ; num%prime == 0; num /= prime {
		}
	}
	return num == 1
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "is ugly num:", isUgly(i))
	}
}

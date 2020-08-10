//https://leetcode.com/problems/perfect-number/description/
package main

import "fmt"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum == num
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d checkPerfectNumber: %v\n", i, checkPerfectNumber(i))
	}
}

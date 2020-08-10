//https://leetcode.com/problems/happy-number/description/

package main

import (
	"fmt"
	"time"
)

func sqrtSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	hashmap := [1000]int{0}
	n = sqrtSum(n)
	for num := 0; hashmap[n] == 0; {
		hashmap[n] = 1
		if n == 1 {
			return true
		}
		if num++; num > 7 {
			return false
		}
		n = sqrtSum(n)
	}
	return false
}

func main() {
	for i := 0; i < 2000; i++ {
		start := time.Now()
		bret := isHappy(i)
		end := time.Now()
		if bret {
			fmt.Printf("%d is HappyNum, time:%d\n", i, end.Sub(start))
		}
		fmt.Printf("%d isn't HappyNum, time:%d\n", i, end.Sub(start))
	}
}

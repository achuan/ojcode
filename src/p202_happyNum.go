package main

import "fmt"

//https://leetcode-cn.com/problems/happy-number/
func isHappy(n int) bool {
	visitMap := make(map[int]bool)
	for n > 0 {
		var sum int
		for n > 0 {
			sum += (n%10)*(n%10)
			n /= 10
		}
		if sum == 1 {
			return true
		}
		if _, ok := visitMap[sum]; !ok{
			visitMap[sum] = true
		} else {
			return false
		}

		n = sum
	}
	return false
}

func main(){
	for i := 2; i <= 20 ; i++{
		fmt.Printf("num %v is perfect %v\n", i, isHappy(i))
	}
}


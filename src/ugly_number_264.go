//https://leetcode.com/problems/ugly-number-ii/description/
package main

import "fmt"

func MIN(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func UglyNumberArray() []int {
	array := []int{1}
	L2, L3, L5 := 0, 0, 0
	for len(array) <= 1690 {
		tmp := MIN(array[L2]*2, MIN(array[L3]*3, array[L5]*5))
		array = append(array, tmp)
		if tmp%2 == 0 {
			L2++
		}
		if tmp%3 == 0 {
			L3++
		}
		if tmp%5 == 0 {
			L5++
		}
	}
	return array
}

func nthUglyNumber(n int) int {
	ugly_array := UglyNumberArray()
	return ugly_array[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}

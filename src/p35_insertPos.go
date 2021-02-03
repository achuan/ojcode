package main

import "fmt"

//https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	var ret int
	for _, item := range nums {
		if item >= target {
			break
		}
		ret++
	}
	return ret
}

func main() {
	testArray := []int{1, 3, 5, 6}
	targetArray := []int{5, 2, 7}

	for _, target := range targetArray {
		fmt.Printf("array %v insert target %v in pos %v\n",
			testArray, target, searchInsert(testArray, target))
	}
}

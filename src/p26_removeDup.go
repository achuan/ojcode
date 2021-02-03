package main

import "fmt"

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == nums[left+1] {
			for i := left + 1; i < right; i++ {
				nums[i] = nums[i+1]
			}
			right--
		} else {
			left++
		}
	}
	return right + 1
}

func main() {
	testArray := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("test of %v result is %v, %v\n", testArray, removeDuplicates(testArray), testArray)
}

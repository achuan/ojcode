package main

import "fmt"

//https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val {
			for i := left; i < right; i++ {
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
	testArray := []int{3, 2, 2, 3}
	len := removeElement(testArray, 3)
	for i := 0; i < len; i++ {
		fmt.Printf("%v ", testArray[i])
	}
}

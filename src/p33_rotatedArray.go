package main

import "fmt"

//https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target{
			return mid
		}
		if nums[0] <= nums[mid]{
			if target >= nums[0] && target < nums[mid]{
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums)-1]{
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main(){
	testArray := [][]int{
		{3, 1},
	}

	for _, item := range testArray{
		for _, key := range item{
			fmt.Printf("array:%v, target:%v, pos:%v\n", item, key, search(item, key))
		}
		fmt.Printf("array:%v, target:%v, pos:%v\n", item, 1000, search(item, 1000))
	}
}

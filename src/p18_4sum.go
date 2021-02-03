package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	var retArray [][]int

	sort.Ints(nums)
	size := len(nums)

	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < size; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, y := j+1, size-1; k < y; {
				if k > j+1 && nums[k] == nums[k-1] {
					k++
					continue
				}
				sum := nums[i] + nums[j] + nums[k] + nums[y]
				if sum < target {
					k++
				} else if sum == target {
					retArray = append(retArray, []int{nums[i], nums[j], nums[k], nums[y]})
					k++
					y--
				} else {
					y--
				}
			}
		}
	}
	return retArray
}

func main() {
	testArray := [][]int{
		{-1, -5, -5, -3, 2, 5, 0, 4},
		{-1, 0, 1, 2, -1, -4},
		{-4, -3, -2, -1, 0, 1, 3, 4, 5},
		{-2, 0, 0, 2, 2},
		{1, 0, -1, 0, -2, 2},
	}
	for _, item := range testArray {
		fmt.Printf("test %v result is :%v\n", item, fourSum(item, -7))
	}
}

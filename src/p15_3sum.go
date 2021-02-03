package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/3sum/

func threeSum(nums []int) [][]int {
	size := len(nums)
	var retArray [][]int

	sort.Ints(nums)

	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, size-1; j < size && k > j; {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				if j > 1 && k < size-1 && nums[j] == nums[j-1] && nums[k] == nums[k+1] {

				} else {
					retArray = append(retArray, []int{nums[i], nums[j], nums[k]})
				}
				j++
				k--

			}
		}
	}
	return retArray
}

func main() {
	testArray := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{-4, -3, -2, -1, 0, 1, 3, 4, 5},
		{-2, 0, 0, 2, 2},
	}
	for _, item := range testArray {
		fmt.Printf("test %v result is :%v\n", item, threeSum(item))
	}
}

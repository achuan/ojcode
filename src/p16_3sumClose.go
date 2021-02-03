package main

import (
	"fmt"
	"math"
	"sort"
)

//https://leetcode-cn.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	size := len(nums)

	var gap int
	for i := 0; i < size; i++ {
		for j, k := i+1, size-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum < target {
				j++
			} else {
				k--
			}
			if gap == 0 || math.Abs(float64(sum-target)) < math.Abs(float64(gap)) {
				gap = sum - target
			}
		}
	}

	return gap + target
}

func main() {
	testArray := [][]int{
		{-2, 0, 0, 2, 2},
		{-1, 0, 1, 2, -1, -4},
		{-4, -3, -2, -1, 0, 1, 3, 4, 5},
	}
	for _, item := range testArray {
		fmt.Printf("test %v result is :%v\n", item, threeSumClosest(item, 1))
	}
}

package main

//https://leetcode-cn.com/problems/container-with-most-water/

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {

	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	testArray := [][]int{{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 8, 6, 2, 5, 4, 8, 3, 5}}
	for _, testHeight := range testArray {
		fmt.Printf("test result of %v is %v\n", testHeight, maxArea(testHeight))
	}
}

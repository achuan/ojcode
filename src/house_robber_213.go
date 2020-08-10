//https://leetcode.com/problems/house-robber-ii/description/
package main

import "fmt"

func rob_house(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mark := 0
	money := make([]int, len(nums))
	for index, value := range nums {
		if index == 0 {
			money[index] = value
		} else if index == 1 {
			if nums[index-1] > nums[index] {
				money[index] = nums[index-1]
			} else {
				money[index] = nums[index]
				mark = index
			}
		} else {
			if mark == index-1 {
				if money[index-2]+nums[index] > money[index-1] {
					money[index] = money[index-2] + nums[index]
					mark = index
				} else {
					money[index] = money[index-1]
				}
			} else {
				money[index] = money[index-1] + nums[index]
				mark = index
			}
		}
	}
	return money[len(money)-1]
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	value1 := rob_house(nums[:len(nums)-1])
	value2 := rob_house(nums[1:])

	if value2 > value1 {
		value1 = value2
	}

	return value1
}

func main() {
	house_array := [][]int{
		{},
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
		{1, 9, 8, 1, 2},
		{1, 9, 8, 3, 2},
	}
	for _, house := range house_array {
		fmt.Println("hourse:", house, "robber:", rob(house))
	}
}

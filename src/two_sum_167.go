//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
package main

import "fmt"

func binarySearch(numbers []int, target int) int {
	left := 0
	right := len(numbers) - 1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers) && numbers[i] <= target/2; i++ {
		right := binarySearch(numbers[i+1:], target-numbers[i])
		if right != -1 {
			return []int{i + 1, right + i + 2}
		}
	}
	return nil
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 18
	param := twoSum(numbers, target)
	fmt.Printf("array:%v\n", numbers)
	fmt.Printf("answer of :%d :%v\n", target, param)
}

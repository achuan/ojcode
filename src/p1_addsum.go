//https://leetcode-cn.com/problems/two-sum/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for index, elem := range nums {
		if prev, ok := numMap[target-elem]; ok {
			return []int{prev, index}
		} else {
			numMap[elem] = index
		}
	}
	return nil
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 17

	sumIndex := twoSum(numbers, target)

	fmt.Printf("array:%v\n", numbers)
	if sumIndex != nil {
		fmt.Printf("answer of %v is :%d, %d\n", target, sumIndex[0], sumIndex[1])
	}
}

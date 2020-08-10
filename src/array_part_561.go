//https://leetcode.com/problems/array-partition-i/description/
package main

import "fmt"

func quickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	quickSort(values[:head])
	quickSort(values[head+1:])
}

func arrayPairSum(nums []int) int {
	sum := 0
	quickSort(nums)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	array_list := [][]int{{1, 2, 3, 2}, {1, 1}, {2, 4, 1, 3}, {1, 2}}
	for _, array := range array_list {
		fmt.Println(array)
		fmt.Println("result: ", arrayPairSum(array))
		fmt.Println(array)
	}
}

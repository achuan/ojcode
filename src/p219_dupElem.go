package main

import "fmt"

//https://leetcode-cn.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	size := len(nums)
	numMap := make(map[int]int)
	for i := 0; i < size; i++{
		if prePos, ok := numMap[nums[i]]; !ok{
			numMap[nums[i]] = i
		} else {
			if i - prePos <= k{
				return true
			}
			numMap[nums[i]] = i
		}
	}
	return false
}

func main(){
	testArray := []int{1,2,3,1,2,3}
	k := 2
	fmt.Print("result of %v is %v\n", testArray, containsNearbyDuplicate(testArray, k))
}

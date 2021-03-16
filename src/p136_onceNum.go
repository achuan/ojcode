package main

import "fmt"

//https://leetcode-cn.com/problems/single-number/

//使用异或处理
func singleNumber(nums []int) int {
	var target int
	for _, item := range nums{
		target ^= item
	}
	return target
}

func main(){
	testArray := [][]int{
		{2,2,1},
		{4,1,2,1,2},
	}
	for _, item := range testArray{
		fmt.Printf("%v, result is:%v\n", item, singleNumber(item))
	}
}


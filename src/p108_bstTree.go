package main

import "fmt"

//https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root TreeNode
	if len(nums) == 0{
		return nil
	}
	size := len(nums)
	root.Val = nums[size/2]
	if size/2 > 0 {
		root.Left = sortedArrayToBST(nums[0:size/2])
	}
	if size/2 < size - 1 {
		root.Right = sortedArrayToBST(nums[size/2+1:])
	}

	return &root
}

func main(){
	nums := []int{-10,-3,0,5,9}
	root := sortedArrayToBST(nums)
	fmt.Printf("%v\n", root)
}
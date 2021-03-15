package main

import "fmt"

//https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {

	var helper func(preSum int, root *TreeNode)int
	helper = func(preSum int, root *TreeNode) int{
		if root == nil{
			return 0
		}

		sum := preSum * 10 + root.Val
		if root.Left == nil && root.Right == nil{
			return sum
		}

		return helper(sum, root.Left) + helper(sum, root.Right)
	}
	return helper(0, root)
}

func createTree(root **TreeNode, i int, array []int){
	size := len(array)
	if i >= size || array[i] == -1 {
		return
	}

	if *root == nil {
		*root = new(TreeNode)
	}

	(*root).Val = array[i]
	createTree(&((*root).Left), 2*i+1, array)
	createTree(&((*root).Right), 2*i+2, array)
}

func initTree(array []int)*TreeNode{
	root := &TreeNode{0, nil, nil}
	createTree(&root, 0, array)
	return root
}

func main(){
	testArray := [][]int{
		{1,1},
		{1, 2, 3},
		{3,1,5,0,2,4,6,-1,-1,-1,3},
		{5,2,7,1,3,6,9,-1,-1,-1,-1,-1,-1,8,10},
		{5,1,4,-1,-1,3,6},
	}

	for _, array := range testArray{
		tree := initTree(array)
		fmt.Printf("%v,  result :%v\n", array, sumNumbers(tree))
	}
}

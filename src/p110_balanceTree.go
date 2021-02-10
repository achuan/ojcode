package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/balanced-binary-tree/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func treeDepth (root *TreeNode)(bool, int) {
	if root == nil {
		return true, 0
	}
	var balance bool
	var leftDepth, rightDepth int
	if root.Left != nil {
		balance, leftDepth = treeDepth(root.Left)
		if !balance {
			return false, 0
		}
	}
	if root.Right != nil {
		balance, rightDepth = treeDepth(root.Right)
		if !balance{
			return false, 0
		}
	}

	if leftDepth >= rightDepth{
		return leftDepth - rightDepth <= 1, leftDepth+1
	}

	return rightDepth - leftDepth <= 1, rightDepth+1
}

func isBalanced(root *TreeNode) bool {
	balance, _ := treeDepth(root)
	return balance
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
		{1,2,2,3,3,-1,-1,4,4},
		{3,9,20,-1,-1,15,7},
		{3,1,5,0,2,4,6,-1,-1,-1,3},
		{1,1},
		{5,2,7,1,3,6,9,-1,-1,-1,-1,-1,-1,8,10},
		{5,1,4,-1,-1,3,6},
	}

	for _, array := range testArray{
		tree := initTree(array)
		fmt.Printf("%v, isbalanced :%v\n", array, isBalanced(tree))
	}
}
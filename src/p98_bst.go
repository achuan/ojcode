package main

import "fmt"

//https://leetcode-cn.com/problems/validate-binary-search-tree/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func bstInfo(root *TreeNode)(bool, int, int){
	var maxCnt, minCnt int = root.Val, root.Val

	if root.Left != nil{
		if root.Left.Val >= root.Val{
			return false, 0, 0
		}

		valid, minLeft, maxLeft := bstInfo(root.Left)
		if !valid || maxLeft >= root.Val{
			return false, 0, 0
		}
		minCnt = minLeft
	}

	if root.Right != nil{
		if root.Right.Val < root.Val{
			return false, 0, 0
		}

		valid, minRight, maxRight := bstInfo(root.Right)
		if !valid || minRight <= root.Val{
			return false, 0, 0
		}
		maxCnt = maxRight
	}

	return true, minCnt, maxCnt
}

func isValidBST(root *TreeNode) bool {
	valid, _, _ := bstInfo(root)
	return valid
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
		{3,1,5,0,2,4,6,-1,-1,-1,3},
		{1,1},
		{5,2,7,1,3,6,9,-1,-1,-1,-1,-1,-1,8,10},
		{5,1,4,-1,-1,3,6},
	}

	for _, array := range testArray{
		tree := initTree(array)
		fmt.Printf("%v, isvalid bst :%v\n", array, isValidBST(tree))
	}
}

package main

import "fmt"

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var retValue []int
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		retValue = append(retValue, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return retValue
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
		{1, -1, 2, -1, -1, 3},
		{3,1,5,0,2,4,6,-1,-1,-1,3},
		{1,1},
		{5,2,7,1,3,6,9,-1,-1,-1,-1,-1,-1,8,10},
		{5,1,4,-1,-1,3,6},
	}

	for _, array := range testArray{
		tree := initTree(array)
		fmt.Printf("%v,  result :%v\n", array, preorderTraversal(tree))
	}
}


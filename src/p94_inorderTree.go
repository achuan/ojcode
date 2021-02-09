package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	var retArray []int
	if root == nil{
		return []int{}
	}
	if root.Left != nil {
		retArray = append(retArray, inorderTraversal(root.Left)...)
	}
	retArray = append(retArray, root.Val)
	if root.Right != nil{
		retArray = append(retArray, inorderTraversal(root.Right)...)
	}

	return retArray
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
		fmt.Printf("%v, inorderTraversal :%v\n", array, inorderTraversal(tree))
	}
}
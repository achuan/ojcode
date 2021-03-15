package main

import "fmt"

//https://leetcode-cn.com/problems/binary-tree-right-side-view/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var retValue []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var tmpQueue []*TreeNode
		for index, node := range queue {
			if len(queue) == index + 1{
				retValue = append(retValue, node.Val)
			}
			if node.Left != nil{
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil{
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = tmpQueue
	}
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
		{1,2,3,-1,5,-1,4},
		{3,1,5,0,2,4,6,-1,-1,-1,3},
		{1,1},
		{5,2,7,1,3,6,9,-1,-1,-1,-1,-1,-1,8,10},
		{5,1,4,-1,-1,3,6},
	}

	for _, array := range testArray{
		tree := initTree(array)
		fmt.Printf("%v, isvalid bst :%v\n", array, rightSideView(tree))
	}
}
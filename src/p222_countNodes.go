package main

import "fmt"

//https://leetcode-cn.com/problems/count-complete-tree-nodes/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}
	sum := 1
	if root.Left != nil {
		sum += countNodes(root.Left)
	}
	if root.Right != nil{
		sum += countNodes(root.Right)
	}
	return sum
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
		{1,2,3,4,5,6},
	}
	for _, item := range testArray{
		tree := initTree(item)
		fmt.Printf("%v, result:%v\n", item, countNodes(tree))
	}
}
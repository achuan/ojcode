package main

import "fmt"

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
	var levelArray [][]int
	if root == nil{
		return [][]int{}
	}
	var helper func(root *TreeNode, depth int)
	helper = func(root *TreeNode, depth int){
		if root == nil{
			return
		}

		if len(levelArray) <= depth{
			levelArray = append(levelArray, []int{})
		}
		levelArray[depth] = append(levelArray[depth], root.Val)
		helper(root.Left, depth+1)
		helper(root.Right, depth+1)
	}
	helper(root, 0)
	return levelArray
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
	{3,9,20,-1,-1,15,7},
	{0,-3,9,10,-1,5,-1},
	}
	for _, array := range testArray{
		tree := initTree(array)
		fmt.Printf("levelorder of %v is %v\n", array, levelOrder(tree))
	}
}
package main

import "fmt"

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++{
		var p []*TreeNode
		var level []int
		for j := 0; j < len(q); j++{
			node := q[j]
			level = append(level, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil{
				p = append(p, node.Right)
			}
		}
		ret = append([][]int{level}, ret...)
		q = p
	}
	return ret
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
		{3,9,20,-1,-1,15,7},
		{0,-3,9,10,-1,5,-1},
	}
	for _, array := range testArray{
		tree := initTree(array)
		fmt.Printf("levelorder of %v is %v\n", array, levelOrderBottom(tree))
	}
}
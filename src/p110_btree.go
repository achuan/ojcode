package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/balanced-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftH := dfs(root.Left)
		if leftH < 0 {
			return -1
		}
		rightH := dfs(root.Right)
		if rightH < 0 {
			return -1
		}
		if math.Abs(float64(leftH - rightH)) > 1{
			return -1
		}

		return int(math.Max(float64(leftH), float64(rightH))) + 1
	}

	return dfs(root) >= 0
}

func createTree(root **TreeNode, i int, array []int) {
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

func initTree(array []int) *TreeNode {
	root := &TreeNode{0, nil, nil}
	createTree(&root, 0, array)
	return root
}

func main(){
	testArray := [][]int{
		{1, 3,4, -1, -1, 2},
		{3, 1, 4, -1, -1, 2, -1},
		{4, 1, 5, 3, -1, -1, -1, -1, 2},
	}
	for _, array := range testArray {
		tree := initTree(array)
		fmt.Printf("%v is isbalanced:%v\n", array, isBalanced(tree))
	}
}
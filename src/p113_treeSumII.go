package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/path-sum-ii/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	var path [][]int

	var dfs func(root *TreeNode, preNode []int, targetSum int)
	dfs = func(root *TreeNode, preNode []int, targetSum int){
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if root.Val == targetSum {
				value := append(append([]int(nil), preNode...), root.Val)
				path = append(path, value)
			}
			return
		}

		preNode = append(preNode, root.Val)
		if root.Left != nil {
			dfs(root.Left, preNode, targetSum - root.Val)
		}
		if root.Right != nil {
			dfs(root.Right, preNode, targetSum - root.Val)
		}
		preNode = preNode[0:len(preNode)-1]
	}
	dfs(root, []int{}, targetSum)
	return path
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
	target := 22
	testArray := []int{5, 4, 8,11,-1,13,4,7,2,-1,-1, -1, -1,5,1}
	tree := initTree(testArray)
	fmt.Printf("target:%v, path:%v\n", target, pathSum(tree, target))
}

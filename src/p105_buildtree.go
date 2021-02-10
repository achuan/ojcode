package main

import "fmt"

//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var root TreeNode
	root.Val = preorder[0]
	for i := 0; i < len(inorder); i++{
		if inorder[i] == root.Val{
			root.Left = buildTree(preorder[1:i+1], inorder[0:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return &root
}

func main(){
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	root := buildTree(preorder, inorder)
	fmt.Printf("%v", root)
}


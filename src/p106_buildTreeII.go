package main

import "fmt"

//https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 ||
		len(postorder) == 0 ||
		len(inorder) != len(postorder){
		return nil
	}
	var root TreeNode
	size := len(postorder)
	root.Val = postorder[size-1]
	for i := 0; i < size; i++{
		if inorder[i] == root.Val{
			root.Left = buildTree(inorder[0:i], postorder[0:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:size-1])
			break
		}
	}
	return &root
}

func main(){
	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	root := buildTree(inorder, postorder)
	fmt.Printf("%v", root)
}

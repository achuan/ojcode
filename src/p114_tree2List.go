package main

//https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil{
		return
	}

	var lastNode *TreeNode
	if root.Right != nil {
		flatten(root.Right)
	}

	if root.Left != nil {
		flatten(root.Left)
		for lastNode = root.Left; lastNode != nil;lastNode = lastNode.Right{
			if lastNode.Right == nil{
				break
			}
		}
		lastNode.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	return
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
		{1,-1,2,-1, -1,3},
		{1,2,5,3,4,-1,6},
		{1,2,5,3,4,10,6},

	}
	for _, array := range testArray{
		tree := initTree(array)
		flatten(tree)
	}
}
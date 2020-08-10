//https://leetcode.com/problems/construct-string-from-binary-tree/description/
package main

import (
	"fmt"
	"strconv"
)

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	tree := string(strconv.Itoa(t.Val))
	if t.Left == nil && t.Right == nil {
		return tree
	}
	tree += "(" + tree2str(t.Left) + ")"
	if t.Right != nil {
		tree += "(" + tree2str(t.Right) + ")"
	}

	return tree
}

func main() {
	tree := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(tree2str(tree))
	tree = &TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(tree2str(tree))
}

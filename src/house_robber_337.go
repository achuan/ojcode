//https://leetcode.com/problems/house-robber-iii/description/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//ret[0]rob this node, ret[1]don't rob this node
func rob_house(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := rob_house(root.Left)
	right := rob_house(root.Right)

	ret := make([]int, 2)
	ret[0] = left[1] + root.Val + right[1]
	ret[1] = max(left[0], left[1]) + max(right[0], right[1])
	return ret
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := rob_house(root)
	return max(ret[0], ret[1])
}

func main() {
	root := &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, &TreeNode{1, nil, nil}}}
	fmt.Println(rob(root))
}

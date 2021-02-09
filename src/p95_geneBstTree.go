package  main

import "fmt"

//https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func numBstComb(begin, end int)[]*TreeNode{
	var treeArray []*TreeNode
	if begin == end {
		treeArray = append(treeArray, &TreeNode{Val: begin})
		return treeArray
	}

	var root TreeNode
	for i := begin; i <= end; i++{
		root.Val = begin
		if i == begin{
			subTree := numBstComb(begin+1, end)
			for _, item := range subTree{
				treeArray = append(treeArray, &TreeNode{Val: i, Left: nil, Right: item})
			}
		} else if i == end{
			subTree := numBstComb(begin, end-1)
			for _, item := range subTree{
				treeArray = append(treeArray, &TreeNode{Val: i, Left: item, Right: nil})
			}
		} else {
			leftArray := numBstComb(begin, i-1)
			rightArray := numBstComb(i+1, end)
			for _, left := range leftArray{
				for _, right := range rightArray{
					treeArray = append(treeArray, &TreeNode{
						Val: i,
						Left: left,
						Right: right,
					})
				}
			}
		}
	}

	return treeArray
}

func generateTrees(n int) []*TreeNode {
	return numBstComb(1, n)
}

func main(){
	for i := 1; i <= 5; i++{
		forest := generateTrees(i)
		fmt.Printf("forest of %v is %v\n", i, len(forest))
	}
}
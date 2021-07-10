package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {

	var sortTree func (head, tail *ListNode) *TreeNode
	sortTree = func(head, tail *ListNode) *TreeNode{
		if head == tail {
			return nil
		}
		slow, fast := head, head
		for fast != tail && fast.Next != tail{
			slow = slow.Next
			fast = fast.Next.Next
		}
		tree := TreeNode{Val: slow.Val}
		tree.Left = sortTree(head, slow)
		tree.Right = sortTree(slow.Next, tail)
		return &tree
	}
	return sortTree(head, nil)
}

func getListNode(array []int) *ListNode {
	headList := &ListNode{0, nil}
	retList := headList
	for _, elem := range array {
		headList.Next = &ListNode{elem, nil}
		headList = headList.Next
	}
	return retList.Next
}

func main(){
	testArray := [][]int{
		{ -10, -3, 0, 5, 9},
		{ -10, -3, 0, 5, 9, 10},
		{ -10, -3, 0, 5, 9, 10, 12},
		{ -10, -3, 0, 5, 9, 10, 12, 15},
	}
	for _, array := range testArray {
		list := getListNode(array)
		tree := sortedListToBST(list)
		fmt.Printf("%v is isbalanced:%v\n", array, tree)
	}
}
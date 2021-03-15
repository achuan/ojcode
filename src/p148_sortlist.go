package main

import "fmt"
//https://leetcode-cn.com/problems/sort-list/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(head1, head2 *ListNode) *ListNode{
	dummyHead := &ListNode{}
	retHead := dummyHead
	for head1 != nil && head2 != nil{
		if head1.Val <= head2.Val{
			dummyHead.Next = head1
			head1 = head1.Next
		} else {
			dummyHead.Next = head2
			head2 = head2.Next
		}
		dummyHead = dummyHead.Next
	}
	if head1 != nil {
		dummyHead.Next = head1
	}
	if head2 != nil {
		dummyHead.Next = head2
	}
	return retHead.Next
}

func sort(head, tail *ListNode) *ListNode{
	if head == nil {
		return head
	}

	if head.Next == tail{
		head.Next = nil
		return head
	}

	fast, slow := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail{
			fast = fast.Next
		}
	}

	mid := slow

	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
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

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println("")
}

func main() {
	testArray := [][]int{
		{4, 2, 1, 3},
		{-1, 5, 3, 4, 0},
		{4,19,14,5,-3,1,8,5,11,15},
	}

	for _, item := range testArray {
		head := getListNode(item)
		printList(sortList(head))
	}
}
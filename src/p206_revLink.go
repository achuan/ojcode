package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

//https://leetcode-cn.com/problems/reverse-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	//curr := head
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
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
	testArray := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5, 6}, {1}, {1, 2}}

	for _, item := range testArray {
		list := getListNode(item)
		printList(list)
		reverseList := reverseList(list)
		printList(reverseList)
	}
}

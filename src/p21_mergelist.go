package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	mergeList := &ListNode{0, nil}
	preHead := mergeList

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			mergeList.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			mergeList.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		mergeList = mergeList.Next
	}
	if l1 != nil {
		mergeList.Next = l1
	}
	if l2 != nil {
		mergeList.Next = l2
	}
	return preHead.Next
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
	//l1 := []int{1,2,4}
	//l2 := []int{1,3,5,8,9}

	l1 := []int{}
	l2 := []int{}

	listOne := getListNode(l1)
	listTwo := getListNode(l2)

	printList(listOne)
	printList(listTwo)
	mergeList := mergeTwoLists(listOne, listTwo)
	printList(mergeList)
}

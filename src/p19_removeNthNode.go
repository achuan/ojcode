package main

import "fmt"

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prevHead := &ListNode{0, head}
	p, q := prevHead, prevHead
	for i := 0; i <= n && q != nil; i++ {
		q = q.Next
	}

	for q != nil {
		q = q.Next
		p = p.Next
	}

	p.Next = p.Next.Next
	return prevHead.Next
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
		{1, 2, 3, 4, 5}, {1}, {1, 2}}

	for _, array := range testArray {
		for i := 1; i <= len(array); i++ {
			listNode := getListNode(array)
			newList := removeNthFromEnd(listNode, i)
			fmt.Printf("remove list:%v, %v-th element, result:\n", array, i)
			printList(newList)
		}
	}
}

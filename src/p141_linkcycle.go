package main

import "fmt"

//https://leetcode-cn.com/problems/linked-list-cycle/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
 	Val int
 	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil{
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			return true
		}
	}
	return false
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

	array := []int{1,2,3,4,5}
	head := getListNode(array)

	head.Next.Next.Next = head.Next

	fmt.Printf("result %v\n", hasCycle(head))
}
package main

import "fmt"

//https://leetcode-cn.com/problems/remove-linked-list-elements/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
  Val int
  Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{0, head}
	if head == nil{
		return head
	}

	prevNode := dummyHead
	for head != nil{
		if head.Val == val{
			prevNode.Next = head.Next
		} else {
			prevNode = head
		}
		head = head.Next
	}

	return dummyHead.Next
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


func main(){
	array := []int{1,2,6,3,4,5,6}
	target := 6

	head := getListNode(array)
	head = removeElements(head, target)
	printList(head)

}
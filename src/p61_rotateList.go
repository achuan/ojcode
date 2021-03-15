package main

import "fmt"

//https://leetcode-cn.com/problems/rotate-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	tmp := head
	size := 0

	if head == nil{
		return nil
	}
	for tmp != nil {
		size++
		tmp = tmp.Next
	}

	k = k%size
	if k == 0{
		return head
	}

	slow, fast := head, head
	for ;k >= 0; k--{
		fast=fast.Next
	}
	father := &ListNode{0, head}
	for fast != nil{
		father = fast
		fast = fast.Next
		slow = slow.Next
	}

	retNode := slow.Next
	slow.Next = father.Next
	father.Next = head

	return retNode
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
	l1 := []int{1, 2}
	k := 1

	head := getListNode(l1)
	printList(rotateRight(head, k))
}

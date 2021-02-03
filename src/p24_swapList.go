package main

import "fmt"

//https://leetcode-cn.com/problems/swap-nodes-in-pairs/

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prevHead := &ListNode{0, head}
	curNode := prevHead

	for curNode != nil && curNode.Next != nil && curNode.Next.Next != nil {
		tmp := curNode.Next.Next
		curNode.Next.Next = tmp.Next
		tmp.Next = curNode.Next
		curNode.Next = tmp
		curNode = curNode.Next.Next
	}

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

func main(){
	testArray := [][]int{
		{1, 2, 3, 4, 5}, {1}, {1, 2}, {1,2,3,4,5,6,7,8}}

	for _, array := range testArray {
		listNode := getListNode(array)
		printList(listNode)
		swapList := swapPairs(listNode)
		printList(swapList)
	}
}
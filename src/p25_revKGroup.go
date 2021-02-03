package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	prevHead := &ListNode{0, head}
	curPrevNode := prevHead
	if k <= 1 {
		return head
	}

	count := k
	for head != nil && head.Next != nil {
		//保证不足k个元素，保持不变
		tail := curPrevNode
		for i := 0; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil{
			break
		}

		//翻转 curPrevNode.next 和 head
		tmp := head.Next
		head.Next = head.Next.Next
		tmp.Next = curPrevNode.Next
		curPrevNode.Next = tmp
		count--
		if count == 1{
			count = k
			curPrevNode = head
			head = head.Next
		}
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
		for i := 1; i <= len(array); i++{
			listNode := getListNode(array)
			fmt.Printf("array:%v, group:%v\n", array, i)
			swapList := reverseKGroup(listNode, i)
			printList(swapList)
		}

	}
}
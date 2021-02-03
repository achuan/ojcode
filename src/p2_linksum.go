//https://leetcode-cn.com/problems/add-two-numbers/
package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	indexList := &ListNode{0, nil}
	preHead := indexList
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		indexList.Next = &ListNode{sum % 10, nil}
		indexList = indexList.Next
	}

	if carry > 0 {
		indexList.Next = &ListNode{carry, nil}
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
	l1 := []int{5, 6, 4}
	l2 := []int{7, 0, 8}

	listOne := getListNode(l1)
	listTwo := getListNode(l2)

	sumList := addTwoNumbers(listOne, listTwo)
	printList(sumList)
}

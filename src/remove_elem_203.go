//https://leetcode.com/problems/remove-linked-list-elements/description/
package main

import "fmt"

/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{-1, head}
	node := prev
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}
	return node.Next
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
}

func main() {
	array := []ListNode{{0, nil}, {1, nil}, {3, nil}, {4, nil}, {1, nil}, {2, nil}, {1, nil}, {4, nil}}
	for i := 0; i < len(array)-1; i++ {
		head := &array[i]
		head.Next = &array[i+1]
	}
	list_node := array[0]
	printList(&list_node)
	removeElements(&list_node, 1)
	printList(&list_node)

}

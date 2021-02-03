package main

import "fmt"

//https://leetcode-cn.com/problems/merge-k-sorted-lists/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
      Val int
      Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	prevHead := &ListNode{0, nil}
	mergeList := prevHead

	if len(lists) == 0{
		return nil
	}

	var curr int
	for {
		for i := 0; i < len(lists); i++{
			if lists[curr] == nil {
				curr = i
				continue
			}

			if lists[i] != nil && lists[i].Val < lists[curr].Val{
				curr = i
			}
		}
		if lists[curr] == nil {
			break
		}
		mergeList.Next = &ListNode{lists[curr].Val, nil}
		mergeList = mergeList.Next
		lists[curr] = lists[curr].Next
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
	testArray := [][]int{{1,1,1},{1,3,4},{2,6}}
	var listArray []*ListNode
	for _, item := range testArray{
		fmt.Printf("array %v\n", item)
		listArray = append(listArray, getListNode(item))
	}

	mergeList := mergeKLists([]*ListNode{})
	printList(mergeList)
}
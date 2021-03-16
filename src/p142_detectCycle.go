package main

//https://leetcode-cn.com/problems/linked-list-cycle-ii/
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for{
		if fast == nil || fast.Next == nil{
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			break
		}
	}
	fast = head
	for fast != slow{
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

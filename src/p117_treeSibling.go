package main

import "fmt"

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil{
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0{
		var tmpQueue []*Node

		size := len(queue)
		for i := 0; i < size; i++{
			if i + 1 < size{
				queue[i].Next = queue[i+1]
			}

			if queue[i].Left != nil{
				tmpQueue = append(tmpQueue, queue[i].Left)
			}
			if queue[i].Right != nil{
				tmpQueue = append(tmpQueue, queue[i].Right)
			}
		}
		queue = tmpQueue
	}
	return root
}

func main(){
	testArray := [][]int{
		{1,2,3,4,5,6,7},
	}
	for _, item := range testArray{
		tree := initTree(item)
		newtree := connect(tree)
		fmt.Printf("%v, %v\n", item, newtree)
	}
}
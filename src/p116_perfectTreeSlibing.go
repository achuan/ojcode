package main

import "fmt"

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
/**
 * Definition for a Node.
 */
type Node struct {
      Val int
      Left *Node
      Right *Node
      Next *Node
}

//我的方法
func connect(root *Node) *Node {
	if root == nil{
		return nil
	}
	if root.Left != nil{
		root.Left.Next = root.Right
		connect(root.Left)
	}
	if root.Right != nil{
		if root.Next == nil{
			root.Right.Next = nil
		} else {
			root.Right.Next = root.Next.Left
		}
		connect(root.Right)
	}

	return root
}

//层次遍历
func connect2(root *Node) *Node {
	queue := []*Node{root}
	for len(queue) > 0{
		tmp := queue
		queue = nil

		for i := 0; i < len(tmp); i++{
			if i + 1 < len(queue){
				tmp[i].Next = tmp[i+1]
			}
			if tmp[i].Left != nil{
				queue = append(queue, tmp[i].Left)
			}
			if tmp[i].Right != nil{
				queue = append(queue, tmp[i].Right)
			}
		}
	}
	return root
}

//leftmost
func connect3(root *Node) *Node{
	if root == nil{
		return nil
	}

	for leftmost := root; leftmost != nil; leftmost = leftmost.Left{
		for node := leftmost; node != nil ; node = node.Next{
			node.Left.Next = node.Right
			if node.Next != nil{
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

func createTree(root **Node, i int, array []int){
	size := len(array)
	if i >= size || array[i] == -1 {
		return
	}

	if *root == nil {
		*root = new(Node)
	}

	(*root).Val = array[i]
	createTree(&((*root).Left), 2*i+1, array)
	createTree(&((*root).Right), 2*i+2, array)
}

func initTree(array []int)*Node{
	root := &Node{0, nil, nil, nil}
	createTree(&root, 0, array)
	return root
}

func main(){
	testArray := [][]int{
		{1,2,3,4,5,6,7},
	}
	for _, item := range testArray{
		tree := initTree(item)
		//newtree := connect(tree)
		newtree := connect2(tree)
		fmt.Printf("%v, %v\n", item, newtree)
	}
}

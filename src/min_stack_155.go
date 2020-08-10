//https://leetcode.com/problems/min-stack/description/
package main

import "fmt"

type MinStack struct {
	elem []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	obj := MinStack{}
	return obj
}

func (this *MinStack) Push(x int) {
	this.elem = append(this.elem, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else if this.min[len(this.min)-1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	this.elem = this.elem[:len(this.elem)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.elem[len(this.elem)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
* Your MinStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	for i := 10; i > 0; i-- {
		obj.Push(i)
		obj.Push(i + 2)
	}
	for i := 0; i < 20; i++ {
		param_3 := obj.Top()
		param_min := obj.GetMin()
		fmt.Printf("i:%d, top value:%d, min:%d\n", i, param_3, param_min)
		obj.Pop()
	}
	//param_4 := obj.GetMin();
}

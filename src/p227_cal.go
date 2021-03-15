package main

import "fmt"

//https://leetcode-cn.com/problems/basic-calculator-ii/
type Stack struct {
	num []int
}

func newStack() *Stack {
	return &Stack{
		num:make([]int, 0),
	}
}

func (s *Stack)push(x int){
	s.num = append(s.num, x)
}

func (s *Stack)pop()int{
	if len(s.num) == 0{
		return 0
	}
	size := len(s.num)
	ret := s.num[size-1]
	s.num = s.num[:size-1]
	return ret
}

func (s *Stack)empty()bool{
	return len(s.num) == 0
}

func isDigit(ch int32)bool{
	if ch >= '0' && ch <= '9'{
		return true
	}
	return false
}

func calculate(s string) int {
	numstack := newStack()
	var num int
	op := '+'
	for index, item := range s {
		if isDigit(item){
			num = 10 * num + int(item - '0')
		}
		if (!isDigit(item) && item != ' ') || index == len(s)-1 {
			if op == '+' {
				numstack.push(num)
			} else if op == '-'{
				numstack.push(-1*num)
			} else if op == '*'{
				numstack.push(num * numstack.pop())
			} else if op == '/'{
				numstack.push(numstack.pop() / num)
			}
			op = item
			num = 0
		}
	}
	var sum int
	for !numstack.empty(){
		sum += numstack.pop()
	}
	return sum
}

func main(){
	testArray := []string{
		"3+2*2",
		"3*2+2",
		"3*2+3*8",
		"3 - 2*2 + 4/2 + 1",
	}
	for _, item := range testArray{
		fmt.Printf("result of %v is %v\n", item, calculate(item))
	}
}
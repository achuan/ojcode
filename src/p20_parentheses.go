package main

import "fmt"

//https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	bracketMap := map[int32]int32{'(': ')',
		'{': '}',
		'[': ']'}

	stack := make([]int32, 0)
	for _, item := range s {
		if bracketMap[item] > 0 {
			stack = append(stack, bracketMap[item])
		} else if len(stack) > 0 && item == stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	testArray := []string{"()", "(])", "()[]{}", "()[]{[()]}", "()[1]{}", "]"}
	for _, item := range testArray {
		fmt.Printf("test of %v result is %v\n", item, isValid(item))
	}
}

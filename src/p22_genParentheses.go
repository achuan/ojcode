package main

import "fmt"

//https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	return getParenthesis("", n, n)
}

func getParenthesis(str string, left, right int)[]string{
	var retArray []string

	if left == 0 && right == 0 {
		return []string{str}
	}

	if left > 0 {
		retArray = append(retArray, getParenthesis(str + "(", left-1, right)...)
	}

	if right > left {
		retArray = append(retArray, getParenthesis(str + ")", left, right-1)...)
	}

	return retArray
}

func main(){
	for i := 1; i < 5; i ++{
		fmt.Printf("n:%v, result:%v\n", i, generateParenthesis(i))
	}
}

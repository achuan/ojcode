package main

import "fmt"

func longestValidParentheses(s string) int {
	size := len(s)
	if size == 0{
		return 0
	}

	maxCnt := 0
	dp := make([]int, size)
	for i := 1; i < size; i++{
		if s[i] == ')' {
			if s[i-1] == '('{
				if i > 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i - dp[i-1] -1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i -dp[i-1] - 2 >= 0{
						dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}

				}
			}
		}

		if maxCnt < dp[i]{
			maxCnt = dp[i]
		}
	}

	return maxCnt
}

func main(){
	testArray := []string{
		"",
		"(()",
		"(())",
		"()(()",
		")()())",
		"(())()",
		"()())((()))",
		")()())()()())(()",
	}
	for _, item := range testArray{
		fmt.Printf("longest parentheses of %v is %v\n", item, longestValidParentheses(item))
	}
}

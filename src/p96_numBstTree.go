package main

import "fmt"

//https://leetcode-cn.com/problems/unique-binary-search-trees/
func numTrees(n int) int {
	summary := make([]int, n+1)
	summary[1] = 1
	for i := 2; i <= n; i++{
		summary[i] = 2*summary[i-1]
		for j := 2; j < i; j++{
			summary[i] += summary[j-1] * summary[i-j]
		}
	}
	return summary[n]
}

func main(){
	for i := 1; i <= 10; i++{
		fmt.Printf("%v bsttree num is %v\n", i, numTrees(i))
	}
}

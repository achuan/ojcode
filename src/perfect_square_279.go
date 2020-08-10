//https://leetcode.com/problems/perfect-squares/description/
package main

import "fmt"

func MIN(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		min := i
		for j := 2; j*j <= i; j++ {
			min = MIN(min, dp[i-j*j]+1)
		}
		dp[i] = min
	}
	return dp[n]
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d numSquares : %d\n", i, numSquares(i))
	}
}

//https://leetcode.com/problems/length-of-last-word/description/
package main

import "fmt"

func lengthOfLastWord(s string) int {
	n := len(s)
	start := -1
	for i := 0; i < n; i++ {
		if s[n-i-1] == ' ' {
			if start >= 0 {
				return i - start
			}
		} else if start == -1 {
			start = i
		}
	}

	if start >= 0 {
		return n - start
	}

	return 0
}

func main() {
	word := "     "
	fmt.Printf("last word lenght of %s is :%d\n", word, lengthOfLastWord(word))
}

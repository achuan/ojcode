package main

import "fmt"

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(str string) int {
	visitMap := map[byte]int{}
	nRetMaxLen, right := 0, -1
	size := len(str)

	for i := 0; i < size; i++ {
		if i != 0 {
			delete(visitMap, str[i-1])
		}
		for right+1 < size && visitMap[str[right+1]] == 0 {
			visitMap[str[right+1]]++
			right++
		}
		if nRetMaxLen < right+1-i {
			nRetMaxLen = right + 1 - i
		}
	}

	return nRetMaxLen
}

func main() {

	inputStr := "pwwkew"

	fmt.Printf("max substr of %s is %d\n", inputStr, lengthOfLongestSubstring(inputStr))
}

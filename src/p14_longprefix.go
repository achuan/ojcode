package main

import "fmt"

//https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var retStr string
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[0][i] != strs[j][i] {
				return retStr
			}
		}
		retStr += string(strs[0][i])
	}
	return retStr
}

func main() {
	testArray := []string{"flower", "flow"}
	fmt.Printf("common prefix is %v\n", longestCommonPrefix(testArray))
}

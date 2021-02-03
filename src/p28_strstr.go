package main

import "fmt"

//https://leetcode-cn.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	target := -1

	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for index := 0; index < len(haystack)-len(needle)+1; {
		if haystack[index] == needle[0] {
			target = index
			for j := 0; j < len(needle); j++ {
				if haystack[index+j] != needle[j] {
					target = -1
					break
				}
			}
			if target != -1 {
				break
			}
		}
		index++
	}
	return target
}

func main() {
	/*
		testArray := []string{"strstr", "abcstr", "abstrcd"}
		for _, item := range testArray{
			fmt.Printf("test of %v is %v\n", item, strStr(item, "str"))
		}*/
	fmt.Printf("test of %v is %v\n", "mississippi", strStr("mississippi", "issipi"))
}

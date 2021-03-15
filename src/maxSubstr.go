package main

import "fmt"

//最长连续子串
// aabbccff -> aabbcc
// adefg -> defg

func maxSubStr(str string, n int)int{
	i, j := 0,1
	var maxLen int

	if len(str) == 0 {
		return 0
	}

	diffMap := make(map[byte]int)
	diffMap[str[0]] = 1

	for j < len(str){
		diff := str[j] - str[j-1]
		if  diff == 0 {
			j++
			continue
		}
		if diff == 1{
			if len(diffMap) < n {
				diffMap[str[j]] = 1
				j++
			} else {
				if maxLen < j - i {
					maxLen = j - i
				}
				i++
				j = i + 1
				diffMap = map[byte]int{}
				diffMap[str[i]] = 1
			}
			continue
		}

		if maxLen < j - i {
			maxLen = j - i
		}
		i = j
		j++
		diffMap = map[byte]int{}
		diffMap[str[i]] = 1
	}

	if maxLen < j - i {
		maxLen = j - i
	}

	return maxLen
}
func main() {
	str := "aaccddddd"
	n := 3
	fmt.Printf("result:%v\n", maxSubStr(str, n));
}

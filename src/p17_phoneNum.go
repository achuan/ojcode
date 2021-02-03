package main

import "fmt"

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	var retArray []string
	if len(digits) == 0 {
		return []string{}
	}

	head := digits[0]
	tailList := letterCombinations(digits[1:])
	for _, item := range phoneMap[string(head)] {
		if len(tailList) == 0 {
			retArray = append(retArray, string(item))
		} else {
			for _, tail := range tailList {
				retArray = append(retArray, string(item)+tail)
			}
		}
	}
	return retArray
}

func main() {
	testArray := []string{
		"29",
		"23",
	}

	for _, item := range testArray {
		fmt.Printf("test of %v, result is:%v\n", item, letterCombinations(item))
	}
}

//https://leetcode.com/problems/add-strings/description/
package main

import "fmt"

import "strconv"

func addStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	var ret string
	value, mark := 0, 0
	for i, j := 0, 0; i < len1 || j < len2; {
		value = mark
		if i < len1 {
			value += int(num1[len1-i-1] - '0')
			i += 1
		}
		if j < len2 {
			value += int(num2[len2-j-1] - '0')
			j += 1
		}

		mark = value / 10
		ret = strconv.Itoa(value%10) + ret
	}

	if mark > 0 {
		ret = strconv.Itoa(mark) + ret
	}

	return ret
}

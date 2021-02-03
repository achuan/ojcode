package main

import "fmt"

//https://leetcode-cn.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	rowIndex := 0
	goingDown := false
	strArr := make([]string, numRows)

	if numRows == 1 {
		return s
	}

	for _, item := range s {
		strArr[rowIndex] += string(item)
		if rowIndex == 0 || rowIndex == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown == true {
			rowIndex += 1
		} else {
			rowIndex -= 1
		}
	}

	var res string
	for _, item := range strArr {
		res += item
	}
	return res
}

func main() {
	testStr := "PAYPALISHIRING"
	fmt.Println("result of %v is %v", testStr, convert(testStr, 3))
	fmt.Println("result of %v is %v", testStr, convert(testStr, 1))
}

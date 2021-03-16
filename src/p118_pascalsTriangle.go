package main

import "fmt"

//https://leetcode-cn.com/problems/pascals-triangle/
//杨辉三角

func generate(numRows int) [][]int {

	retArray := make([][]int, numRows)
	for i := 0; i < numRows; i++{
		retArray[i] = make([]int, i+1)
		retArray[i][0], retArray[i][i] = 1, 1
		for j := 1; j < i; j++{
			retArray[i][j] = retArray[i-1][j-1] + retArray[i-1][j]
		}
	}

	return retArray
}

func main(){
	for i:=1; i <= 5; i++{
		fmt.Printf("value of %v is %v\n", i, generate(i))
	}
}

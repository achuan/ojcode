package main

import "fmt"


// 长度为n的数组，包含0 --（n-1）数，将这个数组重新有序排列好
// 比如: 0,2,1 -> 0,1,2
//    : 0,3,2,4,1 --> 0,1,2,3,4

func order(array []int) []int{
	size := len(array)
	if size == 0{
		return array
	}

	index := 0
	num := 0
	for num < size{
		target := array[index]
		if target == index{
			index++
		} else {
			array[index], array[target] = array[target], array[index]
		}
		num++
	}
	return array
}

func main() {
	testArray := [][]int{
		{8, 9, 7, 1, 2, 4, 3, 6, 5, 0},
	}
	for _, item := range testArray {
		newArray := order(item)
		fmt.Printf("result of %v is %v\n", item, newArray)
	}
}
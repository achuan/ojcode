//https://leetcode.com/problems/base-7/description/
package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(num*-1)
	} else if num == 0 {
		return "0"
	}

	var ret string
	for num/7 > 0 {
		ret = strconv.Itoa(num%7) + ret
		num /= 7
	}
	if num > 0 {
		ret = strconv.Itoa(num%7) + ret
	}
	return ret
}

func main() {
	test_nums := []int{1, 2, 3, 7, 8, 9, 10, 17, 27, -1, -2, -7, -8, -14, -18}
	for _, num := range test_nums {
		fmt.Printf("num:%v, base_7:%v\n", num, convertToBase7(num))
	}
}

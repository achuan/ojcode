//https://leetcode.com/problems/string-compression/description/
package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	var cnt int64 = 1
	comp_str := []byte{chars[0]}
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			cnt += 1
		} else if cnt == 1 {
			comp_str = append(comp_str, chars[i])
		} else {
			comp_str = strconv.AppendInt(comp_str, cnt, 10)
			comp_str = append(comp_str, chars[i])
			cnt = 1
		}
	}
	if cnt > 1 {
		comp_str = strconv.AppendInt(comp_str, cnt, 10)
	}

	chars = append(chars[:0], comp_str...)

	return len(comp_str)
}

func printArray(chars []byte) {
	for _, key := range chars {
		fmt.Printf("%c ", key)
	}
	fmt.Println()
}

func main() {
	str_list := [][]byte{
		{'a', 'b', 'c', 'c'},
		{'a', 'b', 'c'},
		{'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		{'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c'},
	}
	for _, str := range str_list {
		printArray(str)
		num := compress(str)
		printArray(str)
		fmt.Println("array size:", num)
	}
}

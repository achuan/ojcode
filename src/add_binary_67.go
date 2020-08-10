//https://leetcode.com/problems/add-binary/description/
package main

import "fmt"

func addBinary(a string, b string) string {
	var ret string
	mark := 0
	len_a := len(a)
	len_b := len(b)
	minlen := len_a
	if minlen > len_b {
		minlen = len_b
	}
	for i := 0; i < minlen; i++ {
		mark += int(a[len_a-i-1]-'0') + int(b[len_b-i-1]-'0')
		ret = string(mark%2+'0') + ret
		mark /= 2
	}

	for i := minlen; i < len_a; i++ {
		mark += int(a[len_a-i-1] - '0')
		ret = string(mark%2+'0') + ret
		mark /= 2
	}
	for i := minlen; i < len_b; i++ {
		mark += int(b[len_b-i-1] - '0')
		ret = string(mark%2+'0') + ret
		mark /= 2
	}
	if mark > 0 {
		ret = string(mark%2+'0') + ret
	}
	return ret
}

func main() {
	a := "10110"
	b := "101"
	fmt.Printf("%v + %v = %v\n", a, b, addBinary(a, b))
}

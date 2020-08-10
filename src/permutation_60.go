//https://leetcode.com/problems/permutation-sequence/description/
package main

import "fmt"

func initPermutaion(num int) []int {
	array_seq := make([]int, num)
	array_seq[0] = 1
	for i := 1; i < num; i++ {
		array_seq[i] = (i + 1) * array_seq[i-1]
	}
	return array_seq
}

func getPermutation(n int, k int) string {
	var permut_str string
	bits_map := make([]int, n)

	seq_list := initPermutaion(n)
	for index := n - 1; index >= 1 && k > 0; index-- {
		pos := (k + seq_list[index-1] - 1) / seq_list[index-1]
		k -= (pos - 1) * seq_list[index-1]
		for j := 0; j < n; j++ {
			if bits_map[j] == 0 {
				pos--
			}
			if pos == 0 {
				bits_map[j] = 1
				permut_str += string(j + 1 + '0')
				break
			}
		}
	}

	for j := 0; j < n; j++ {
		if bits_map[j] != 1 {
			permut_str += string(j + 1 + '0')
		}
	}
	return permut_str
}

func main() {
	seq_list := initPermutaion(5)
	for n := 1; n < 5; n++ {
		for k := 1; k <= seq_list[n-1]; k++ {
			fmt.Printf("get permutation %v, %v is:%s\n", n, k, getPermutation(n, k))
		}
	}
}

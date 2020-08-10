//https://leetcode.com/problems/1-bit-and-2-bit-characters/description/
package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	if bits[len(bits)-1] != 0 {
		return false
	}

	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else if bits[i] == 0 {
			i++
		}
	}
	if i == len(bits)-1 {
		return true
	}
	return false
}

func main() {
	bits := []int{1, 1, 1, 0, 0}
	fmt.Printf("%v, isOneBit:%v\n", bits, isOneBitCharacter(bits))
}

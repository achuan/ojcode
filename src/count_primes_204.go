//https://leetcode.com/problems/count-primes/description/
package main

import "fmt"

func countPrimes(n int) int {
	isPrime := make([]int, n)
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] != 0 {
			continue
		}
		count++
		for j := i; j*i < n; j++ {
			isPrime[j*i] = 1
		}
	}
	return count
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, ":prime number: ", countPrimes(i))
	}
}

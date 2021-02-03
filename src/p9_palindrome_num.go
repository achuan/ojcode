package main

import "fmt"

//https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var strValue []byte
	for x != 0 {
		strValue = append(strValue, byte(x%10+'0'))
		x = x / 10
	}
	size := len(strValue)
	for i := 0; i < size/2; i++ {
		if strValue[i] != strValue[size-i-1] {
			return false
		}
	}
	return true
}

func main() {

	testArray := []int{
		-121,
		123,
		0,
		2,
		121,
		12345,
		123456789,
		999888999,
	}

	for _, item := range testArray {
		fmt.Printf("result of %v is %v\n", item, isPalindrome(item))
	}

}

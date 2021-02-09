package main

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/multiply-strings/
func multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	var carry int
	mulint := make([]int, len1+len2)
	for i := len1-1; i >= 0; i--{
		for j := len2-1; j >= 0; j--{
			sum := mulint[len1 + len2 - 2 - i - j]  + carry + int((num1[i] - '0') * (num2[j] - '0'))
			carry = sum / 10
			mulint[len1 + len2 - 2 - i - j] = sum % 10
		}
		if carry > 0 {
			mulint[len1 + len2 - 1 - i ] += carry
			carry = 0
		}
	}

	var retstr string
	notZero := false
	for i := len(mulint)-1; i >= 0; i--{
		if notZero == false && mulint[i] == 0{
			continue
		}
		notZero = true
		retstr += string(mulint[i] + '0')
	}

	if len(retstr) == 0 {
		retstr = "0"
	}

	return retstr
}

func main(){
	array1 := []string{"123","1020","456", "789", "999", "000" }
	array2 := []string{"999","1", "10", "23","1230", "999", "0"}

	for _, item1 := range array1{
		for _, item2 := range array2{

			result := multiply(item1, item2)
			fmt.Printf("multipy %v * %v == %v\n", item1, item2, result)

			num1 , _ := strconv.Atoi(item1)
			num2, _ := strconv.Atoi(item2)

			resultInt, _ := strconv.Atoi(result)
			if resultInt != num1 * num2{
				fmt.Printf("error....error....%v != %v*%v\n", resultInt, num1, num2)
			}
		}
	}

}

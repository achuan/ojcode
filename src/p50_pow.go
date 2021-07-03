package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/powx-n/

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	if n < 0 {
		return myPow(1/x, -n)
	}

	powRet := x
	extra := 1.0
	for n / 2 >= 1 {
		if n % 2 == 1{
			extra = extra * powRet
		}
		powRet = powRet * powRet
		n = n / 2
	}

	return extra * powRet
}

func main(){
	testArray := []struct {
		base float64
		n int
	}{
		{2.0000, 10},{2.0000, 9} , {2.1000, 3},{2.00000, -2},
	}

	for _, item := range testArray{
		fmt.Printf("pow of %f %d = %f\n", item.base, item.n, myPow(item.base, item.n))
	}
}
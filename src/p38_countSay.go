package main

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/count-and-say/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prevStr := countAndSay(n - 1)
	var targetCh int32
	var count int
	var retStr string
	for i, item := range prevStr {
		if i == 0 {
			targetCh = item
			count = 1
			continue
		}
		if targetCh == item {
			count++
		} else {
			retStr += strconv.Itoa(count) + string(targetCh)
			targetCh = item
			count = 1
		}
	}
	if count > 0 {
		retStr += strconv.Itoa(count) + string(targetCh)
	}

	return retStr
}

func main() {
	for i := 1; i < 18; i++ {
		fmt.Printf("value:%v countandsay:%v\n", i, countAndSay(i))
	}
}

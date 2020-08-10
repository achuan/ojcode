//https://leetcode.com/problems/judge-route-circle/description/
package main

import "fmt"

func judgeCircle(moves string) bool {
	var hor, ver int
	for _, v := range moves {
		switch {
		case v == 'U', v == 'u':
			ver += 1
		case v == 'D', v == 'd':
			ver -= 1
		case v == 'L', v == 'l':
			hor += 1
		case v == 'R', v == 'r':
			hor -= 1
		}
	}
	return hor == 0 && ver == 0
}

func main() {
	fmt.Println(judgeCircle("UUDD"))
}

package main

import "fmt"

func main() {
	fmt.Println(minFlips(1,2,3))
}

func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 32; i ++ {
		tc := c & 1
		ta := a & 1
		tb := b & 1
		if tc == 0 {
			if ta == 1 {
				res ++
			}
			if tb == 1 {
				res ++
			}
		}else {
			if ta == 0 && tb == 0 {
				res ++
			}
		}
		a = a >> 1
		b = b >> 1
		c = c >> 1
	}
	return res
}

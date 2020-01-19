package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(699))
}

func maximum69Number (num int) int {
	s := strconv.Itoa(num)
	index := -1
	for v, k := range s {
		if k != '9' {
			index = v
			break
		}
	}
	if index == -1 {
		return num
	}
	index = len(s) - index
	extra := 3
	for i := 1; i < index; i ++ {
		extra *= 10
	}
	return extra + num
}


package main

import "fmt"

/**
x的平方根
 */
func main(){
	x := 16
	fmt.Println(mySqrt(x))
}
func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	guess := x / 2
	for !(guess * guess <= x && (guess + 1) * (guess + 1) > x) {
		guess = (guess + x / guess) / 2
	}
	return guess
}


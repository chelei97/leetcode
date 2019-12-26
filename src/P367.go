package main

import "fmt"

func main() {
	n := 16
	fmt.Println(isPerfectSquare(n))
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	guess := num / 2
	for !(guess * guess <= num && (guess + 1) * (guess + 1) > num) {
		guess = (guess + num / guess) / 2
	}
	return guess * guess == num
}

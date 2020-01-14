package main

import "fmt"

func main() {
	fmt.Println(getNoZeroIntegers(1010))
}
func getNoZeroIntegers(n int) []int {
	var res []int
	for i := 1; i <= n / 2; i ++ {
		if !containsZero(i) && !containsZero(n - i) {
			res = append(res, i)
			res = append(res, n - i)
			return res
		}
	}
	return res
}

func containsZero(n int) bool {
	for n > 0 {
		tmp := n % 10
		if tmp == 0 {
			return true
		}
		n = (n - tmp) / 10
	}
	return false
}

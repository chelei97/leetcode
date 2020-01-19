package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1,22))
}

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i ++ {
		if i < 10 {
			res = append(res, i)
			continue
		}
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDividing(n int) bool {
	t := n
	for n > 0 {
		tmp := n % 10
		if tmp == 0 {
			return false
		}
		if t % tmp != 0 {
			return false
		}
		n = (n - tmp) / 10
	}
	return true
}

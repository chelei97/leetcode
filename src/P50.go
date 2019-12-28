package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(-13.62608, 3))
}

func myPow(x float64, n int) float64 {
	res := 1.0
	if n == 0 {
		return res
	}
	res = x
	rec := res
	flag, negative := false, false
	if x < 0 && n & 1 != 0 {
		negative = true
	}
	if n < 0 {
		flag = !flag
		n = -n
	}
	count := 1
	times := 1
	for count != n {
		if count + times <= n {
			res = res * rec
			rec = res
			count = count + times
			times = times * 2
		}else {
			for count + times > n {
				rec = math.Sqrt(rec)
				times = times / 2
			}
			res = res * rec
			count = count + times
		}
	}
	if negative {
		res = -res
	}
	if flag {
		return 1 / res
	}
	return res
}

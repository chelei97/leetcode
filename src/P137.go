package main

import "fmt"

func main() {
	nums := []int{-2,-2,1,1,-3,1,-3,-3,-4,-2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	a,b := 0,0
	for _,v := range nums {
		a = (a^v) & ^b
		b = (b^v) & ^a
	}
	return a
}

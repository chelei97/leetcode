package main

import "fmt"

func main() {
	nums := []int{1,1,2,2,3,5}
	fmt.Println(singleNumber2(nums))
}

func singleNumber2(nums []int) []int {
	var res, tmp int
	for _, k := range nums {
		res = res ^ k
	}
	diff := res & (-res)
	for _, k := range nums {
		if k & diff != 0 {
			tmp = tmp ^ k
		}
	}
	result := make([]int, 2)
	result[0], result[1] = tmp, res ^ tmp
	return result
}

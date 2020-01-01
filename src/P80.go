package main

import "fmt"

func main() {
	nums := []int{0,0,1,1,1,1,2,3,3}
	k := removeDuplicates(nums)
	fmt.Println(k, nums[ : k])
}

func removeDuplicates(nums []int) int {
	size, length := 0, len(nums)
	for i, j := 0, 0; i < length; {
		for j < length && nums[j] == nums[i] {
			j ++
		}
		tmp := j - i
		if tmp >= 2 {
			tmp = 2
		}
		for n := 0; n < tmp; n ++ {
			nums[size + n] = nums[j - 1]
		}
		size = size + tmp
		i = j
	}
	return size
}

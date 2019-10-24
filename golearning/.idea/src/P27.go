package main

import "fmt"

func main(){
	var nums = [] int {0,4,4,0,4,4,4,0,2}
	var remove int = 4
	var result int = removeElement(nums, remove)
	fmt.Println(result)
	for i := 0; i < result; i++ {
		fmt.Print(nums[i])
	}
}

func removeElement(nums []int, val int) int {
	var left, right, record int = 0, len(nums) - 1, 0

	if left == right {
		if nums[left] == val {
			return 0
		}
	}

	for ; left < right;  {
		if nums[left] == val && nums[right] != val {
			nums[left] = nums[right]
			left ++
			right --
			record ++
		}else if nums[left] == val && nums[right] == val {
			for right >= left && nums[right] == val {
				right --
				record ++
			}
			if right <= 0 {
				return len(nums) - record
			}
			if right >= left {
				nums[left] = nums[right]
				left ++
				right --
				record ++
			}
		}else if nums[left] != val && nums[right] == val {
			left ++
			right --
			record ++
		}else if nums[left] != val && nums[right] != val {
			left ++
		}
	}
	if left == right && nums[left] == val {
		record ++
	}
	return len(nums) - record
}

package main

import "fmt"

func main() {
	answers := []int{1,1,0,0,0}
	fmt.Println(numRabbits(answers))
}

func numRabbits(answers []int) int {
	res := 0
	if len(answers) <= 0 {
		return res
	}
	record := make(map[int]int)
	for i := 0; i < len(answers); i ++ {
		record[answers[i]] ++
	}
	for k, v := range record {
		if k == 0 {
			res = res + v
		}else {
			tmp := v / (k + 1)
			if v % (k + 1) != 0 {
				tmp ++
			}
			res = tmp * (k + 1) + res
		}
	}
	return res
}
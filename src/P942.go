package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID"))
}

func diStringMatch(S string) []int {
	res := make([]int, len(S) + 1)
	max, min := len(S), 0
	for i := 0; i < len(S); i ++ {
		if S[i] == 'I' {
			res[i] = min
			min ++
		}else {
			res[i] = max
			max --
		}
	}
	res[len(S)] = min
	return res
}

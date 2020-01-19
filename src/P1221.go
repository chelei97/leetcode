package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	numR, res := 0, 0
	if s[0] == 'R' {
		numR ++
	}else {
		numR --
	}
	for i := 1; i < len(s); i ++ {
		if s[i] == 'R' {
			numR ++
		}else {
			numR --
		}
		if numR == 0 {
			res ++
		}
	}
	return res
}

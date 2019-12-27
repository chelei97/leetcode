package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) int {
	result := 0
	for i := 0; i < len(s); i ++ {
		result = result + expand1(s, i, i) + expand1(s, i, i + 1)
	}
	return result
}

func expand1(s string, left, right int) int {
	result := 0
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left --
			right ++
			result ++
		}else {
			break
		}
	}
	return result
}
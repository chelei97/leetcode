package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	left, right, match := 0, 0, 0
	start, minLen := 0, len(s) + 1
	record := make(map[byte]int, 0)
	result := make(map[byte]int, 0)
	for _, v := range t {
		record[byte(v)]++
	}
	for right < len(s) {
		if _, ok := record[s[right]]; ok {
			result[s[right]] ++
			if result[s[right]] == record[s[right]] {
				match ++
			}
		}
		right ++

		for match == len(record) {
			if right - left < minLen {
				start = left
				minLen = right - left
			}
			if _, ok := record[s[left]]; ok {
				result[s[left]] --
				if result[s[left]] < record[s[left]] {
					match --
				}
			}
			left ++
		}
	}
	if minLen > len(s) {
		return ""
	}
	return s[start:start+minLen]
}

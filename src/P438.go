package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 {
		return nil
	}
	res := make([]bool, len(s) - len(p) + 1)
	left, right, match := 0, len(p) - 1, 0
	record := make(map[byte]int)
	result := make(map[byte]int)
	for i := 0; i < len(p); i ++ {
		record[p[i]] ++
		result[s[i]] ++
	}
	res[0] = true
	for k, v := range record {
		if result[k] != v {
			res[0] = false
		}else {
			match ++
		}
	}

	for right < len(s) - 1 {
		right ++
		result[s[right]] ++
		if _, ok := record[s[right]]; ok {
			if result[s[right]] == record[s[right]] {
				match ++
			}else if result[s[right]] == record[s[right]] + 1 {
				match --
			}
		}
		result[s[left]] --
		if _, ok := record[s[left]]; ok {
			if result[s[left]] == record[s[left]] {
				match++
			} else if result[s[left]] == record[s[left]] - 1 {
				match--
			}
		}
		left ++

		if match == len(record) {
			res[left] = true
		}
	}

	tmp := make([]int, 0)
	for i := 0; i < len(res); i ++ {
		if res[i] {
			tmp = append(tmp, i)
		}
	}
	return tmp
}

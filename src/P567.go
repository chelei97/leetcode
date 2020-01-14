package main

import (
	"fmt"
)

func main() {
	s1, s2 := "ab", "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	rec := make(map[byte]int)
	lens1, lens2 := len(s1), len(s2)
	for i := 0; i < lens1; i ++ {
		rec[s1[i]] ++
		rec[s2[i]] --
	}
	for k, v := range rec {
		if v == 0 {
			delete(rec, k)
		}
	}
	for i := lens1; len(rec) != 0 && i < lens2; i ++ {
		rec[s2[i]] --
		rec[s2[i - lens1]] ++
		if rec[s2[i]] == 0 {
			delete(rec, s2[i])
		}
		if rec[s2[i - lens1]] == 0 {
			delete(rec, s2[i - lens1])
		}
	}
	if len(rec) == 0 {
		return true
	}
	return false
}
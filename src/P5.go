package main

import (
	"fmt"
)

/**
最长回文子串
 */
func main(){
	s := "abb"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i ++ {
		len1, len2 := expand(s, i, i), expand(s, i, i + 1)
		if len1 < len2 {
			len1 = len2
		}
		if len1 > end - start {
			start = i - (len1 - 1) / 2
			end = i + len1 / 2
		}
	}
	return s[start : end + 1]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right]{
		left --
		right ++
	}
	return right - left - 1
}
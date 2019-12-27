package main

import "fmt"

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	if isInDict(s, wordDict) {
		return true
	}
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i ++ {
		dp[i] = isInDict(s[ : i + 1], wordDict)
	}
	//fmt.Println(dp)
	for i := 0; i < len(s); i ++ {
		if dp[i] {
			continue
		}
		for j := 0; j <= i; j ++ {
			dp[i] = dp[j] && isInDict(s[j + 1 : i + 1], wordDict)
			if dp[i] {
				break
			}
		}
	}
	//fmt.Println(dp)
	return dp[len(s) - 1]
}

func isInDict(s string, wordDict []string) bool {
	for _, v := range wordDict {
		if v == s {
			return true
		}
	}
	return false
}
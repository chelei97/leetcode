package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	lens1, lens2 := len(text1), len(text2)
	var dp [][]int
	for i := 0; i < lens1 + 1; i ++ {
		tmp := make([]int, lens2 + 1)
		dp = append(dp, tmp)
	}
	//dp[i][j]表示text1前i个和text2前j个的LCS
	for i := 1; i < len(dp); i ++ {
		for j := 1; j < len(dp[0]); j ++ {
			if text1[i - 1] == text2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}else {
				dp[i][j] = max10(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[lens1][lens2]
}

func max10(a, b int) int {
	if a > b {
		return a
	}
	return b
}
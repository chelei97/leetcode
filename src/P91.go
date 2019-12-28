package main

import "fmt"

func main() {
	s := "110"
	fmt.Println(numDecodings(s))
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp_0 := 1
	dp_1 := 1
	for i := 1; i < len(s); i ++ {
		tmp := 0
		if s[i] != '0' {
			tmp = tmp + dp_1
		}else if s[i - 1] != '1' && s[i - 1] != '2' {
			return 0
		}else {
			tmp = tmp + dp_0
			dp_0, dp_1 = dp_1, tmp
			continue
		}
		if s[i - 1 : i + 1] < "27" && s[i - 1 : i + 1] > "10" {
			tmp = tmp + dp_0
		}
		dp_0, dp_1 = dp_1, tmp
	}
	return dp_1
}

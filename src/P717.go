package main

import "fmt"

func main() {
	bits := []int{1, 1, 1, 0}
	fmt.Println(isOneBitCharacter(bits))
}

func isOneBitCharacter(bits []int) bool {
	length := len(bits)
	if length < 2 {
		return true
	}
	if bits[length - 2] == 0 {
		return true
	}
	if length == 2 {
		return false
	}
	dp := make([]bool, length - 2)
	for i := 0; i < len(dp); i ++ {
		if bits[i] == 0 {
			dp[i] = true
		}else {
			if i == 0 {
				dp[i] = false
			}else {
				if bits[i - 1] == 1 && dp[i - 1] == false {
					dp[i] = true
				}else {
					dp[i] = false
				}
			}
		}
	}
	fmt.Println(dp[0])
	return !(dp[length - 3])
}

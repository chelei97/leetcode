package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "0000"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	s1 := []byte(s)
	//ith指第几区，jth指第几个字符
	helper4(0, 0, s1, []byte{}, &res)
	return res
}

func helper4(ith, jth int, s, cur []byte, res *[]string) {
	if ith == 4 && jth == len(s) {
		tmp := make([]byte, len(cur))
		tmp = tmp[:len(tmp)-1]
		copy(tmp, cur)
		*res = append(*res, string(tmp))
		return
	}
	if ith < 4 && jth < len(s) {
		if s[jth] == '0' {
			cur = append(cur, s[jth], '.')
			helper4(ith + 1, jth + 1, s, cur, res)
			return
		}
		for i := 0; i <= 2; i ++ {
			for j := 0; jth + j < len(s) && j <= i; j ++ {
				cur = append(cur, s[jth + j])
			}
			cur = append(cur, '.')
			tmp, _ := strconv.Atoi(string(cur[len(cur) - i - 2 : len(cur) - 1]))
			if tmp < 256 {
				helper4(ith + 1, jth + i + 1, s, cur, res)
			}
			cur = cur[ : len(cur) - i - 2]
		}
	}else {
		return
	}
}




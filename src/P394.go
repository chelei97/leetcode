package main

import (
	"container/list"
	"fmt"
)

func main() {
	s := "100[leetcode]"
	fmt.Println(decodeString(s))
}

func decodeString(s string) string {
	multi := 0
	res := make([]byte, 0)
	stack_multi := list.New()
	stack_res := list.New()
	for i := 0; i < len(s); i ++ {
		if s[i] <= '9' && s[i] >= '0' {
			multi = multi * 10 + int(s[i] - '0')
		}else if s[i] == '[' {
			tmp := make([]byte, len(res))
			copy(tmp, res)
			stack_res.PushBack(tmp)
			stack_multi.PushBack(multi)
			res = res[:0]
			multi = 0
		}else if s[i] == ']' {
			tmp := make([]byte, 0)
			tmp_multi := stack_multi.Back().Value.(int)
			stack_multi.Remove(stack_multi.Back())
			for i := 0; i < tmp_multi; i ++ {
				for j := 0; j < len(res); j ++ {
					tmp = append(tmp, res[j])
				}
			}
			res = append(stack_res.Back().Value.([]byte), tmp...)
			stack_res.Remove(stack_res.Back())
		}else {
			res = append(res, s[i])
		}
	}
	return string(res)
}

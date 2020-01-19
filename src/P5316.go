package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "TO BE OR NOT TO BE"
	fmt.Println(printVertically(s))
}

func printVertically(s string) []string {
	record := strings.Split(s, " ")
	rec := make([][]byte, 0)
	for i := 0; i < len(record); i ++ {
		tmp := []byte(record[i])
		rec = append(rec, tmp)
	}
	max := len(rec[0])
	for i := 1; i < len(rec); i ++ {
		if len(rec[i]) > max {
			max = len(rec[i])
		}
	}
	res := make([][]byte, max)
	for i := 0; i < max; i ++ {
		for j := 0; j < len(rec); j ++ {
			if len(rec[j]) > i {
				res[i] = append(res[i], rec[j][i])
			}else {
				res[i] = append(res[i], ' ')
			}
		}
	}
	result := make([]string, max)
	for i := 0; i < max; i ++ {
		j := len(res[i]) - 1
		for res[i][j] == ' ' {
			j --
		}
		if j == len(res[i]) - 1 {
			continue
		}else {
			res[i] = res[i][ : j + 1]
		}
	}
	for i := 0; i < max; i ++ {
		result[i] = string(res[i])
		fmt.Println(len(result[i]))
	}
	return result
}

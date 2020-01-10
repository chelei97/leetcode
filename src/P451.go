package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "leetcode"
	fmt.Println(frequencySort(s))
}

type pair struct {
	value byte
	times int
}

type p451 []pair

func frequencySort(s string) string {
	record := make(map[byte]int)
	for i := 0; i < len(s); i ++ {
		record[s[i]] ++
	}
	var rec p451
	for k, v := range record {
		rec = append(rec, pair{value:k, times:v})
	}
	sort.Slice(rec, func(i, j int) bool {
		return rec[i].times < rec[j].times
	})
	var result []byte
	for i := len(rec) - 1; i >= 0; i -- {
		for j := 0; j < rec[i].times; j ++ {
			result = append(result, rec[i].value)
		}
	}
	return string(result)
}

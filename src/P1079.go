package main

import (
	"fmt"
)

func main() {
	s := "AAABBC"
	fmt.Println(numTilePossibilities(s))
}

var res int
var record2 [26]int
func numTilePossibilities(tiles string) int {
	for i := 0; i < len(tiles); i ++ {
		record2[int(tiles[i] - 'A')] ++
	}
	dfs4()
	return res
}

func dfs4() {
	for i := 0; i < len(record2); i ++ {
		if record2[i] == 0 {
			continue
		}
		record2[i] --
		res ++
		dfs4()
		record2[i] ++
	}
}


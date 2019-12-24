package main

import "fmt"

func main(){
	board := [][]byte{
		{'A','B','C','E'},
	{'S','F','E','S'},
	{'A','D','E','E'},
	}
	fmt.Println(exist(board, "ABCESEEEFSAD"))
}

func exist(board [][]byte, word string) bool {
	record := make([][]int, 0)
	for i := 0; i < len(board); i ++ {
		tmp := make([]int, len(board[0]))
		record = append(record, tmp)
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j ++ {
			if dfs(board, record, i, j, word) {
				return true
			}
			for m := 0; m < len(board); m ++ {
				for n := 0; n < len(board[0]); n ++ {
					record[m][n] = 0
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, record [][]int, i, j int, word string) bool {
	if record[i][j] == 1 {
		return false
	}
	if len(word) == 0 {
		return true
	}
	words := []byte(word)
	if board[i][j] == words[0] {
		if len(words) == 1 {
			return true
		}
		record[i][j] = 1
		if i - 1 >= 0 {
			if dfs(board, record, i - 1, j, string(words[1 : ])) {
				return true
			}
		}
		if i + 1 < len(board){
			if dfs(board, record, i + 1, j, string(words[1 : ])) {
				return true
			}
		}
		if j - 1 >= 0 {
			if dfs(board, record, i, j - 1, string(words[1 : ])) {
				return true
			}
		}
		if j + 1 < len(board[0]){
			if dfs(board, record, i, j + 1, string(words[1 : ])) {
				return true
			}
		}
	}
	record[i][j] = 0
	return false
}

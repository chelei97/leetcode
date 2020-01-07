package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'X', 'O', 'X', 'X', 'O', 'O', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O'},
		{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
	}
	solve(board)
	for i := 0; i < len(board); i ++ {
		str := string(board[i])
		fmt.Println(str)
	}
}

func solve(board [][]byte)  {
	if len(board) < 2 || len(board[0]) < 2 {
		return
	}
	for i := 1; i < len(board) - 1; i ++ {
		for j := 1; j < len(board[0]) - 1; j ++ {
			record := make([][]int, len(board))
			for k := 0; k < len(board); k ++ {
				tmp := make([]int, len(board[0]))
				record[k] = tmp
			}
			if board[i][j] == 'O' && record[i][j] != 1 {
				if dfs3(board, &record, i, j) {
					for m := 1; m < len(board) - 1; m ++ {
						for n := 1; n < len(board[0]) - 1; n ++ {
							if record[m][n] == 1 {
								board[m][n] = 'X'
								record[m][n] = 0
							}
						}
					}
				}
			}
		}
	}
}

func dfs3(board [][]byte, record *[][]int, i, j int) bool {
	if i == 0 || i == len(board) - 1 {
		return false
	}
	if j == 0 || j == len(board[0]) - 1 {
		return false
	}
	(*record)[i][j] = 1
	f1, f2, f3, f4 := true, true, true, true
	if board[i - 1][j] == 'O' && (*record)[i - 1][j] != 1 {
		f1 = dfs3(board, record, i - 1, j)
	}
	if board[i + 1][j] == 'O' && (*record)[i + 1][j] != 1 {
		f2 = dfs3(board, record, i + 1, j)
	}
	if board[i][j - 1] == 'O' && (*record)[i][j - 1] != 1 {
		f3 = dfs3(board, record, i, j - 1)
	}
	if board[i][j + 1] == 'O' && (*record)[i][j + 1] != 1 {
		f4 = dfs3(board, record, i, j + 1)
	}
	return f1 && f2 && f3 && f4
}

package main

import "fmt"

func main() {
	board := [][]int{
		{0,1,0},
		{0,0,1},
		{1,1,1},
		{0,0,0},
	}
	gameOfLife(board)
	fmt.Println(board)
}

func gameOfLife(board [][]int)  {
	if len(board) == 0 {
		return
	}
	row, col := len(board), len(board[0])
	for i := 0; i < row; i ++ {
		for j := 0; j < col; j ++ {
			tmp := count(board, i, j)
			if tmp < 2 && board[i][j] == 1 {
				board[i][j] = 2
			}else if (tmp == 3 || tmp == 2) && board[i][j] == 1 {
				continue
			}else if tmp > 3 && board[i][j] == 1 {
				board[i][j] = 2
			}else if tmp == 3 && board[i][j] == 0 {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < row; i ++ {
		for j := 0; j < col; j ++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			}else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
	return
}

//1是活，2是之前活现在死，0是死，3是之前是现在活
func count(board [][]int, i, j int) int {
	res := 0
	isAlive(i-1, j-1, &res, board)
	isAlive(i, j-1, &res, board)
	isAlive(i-1, j, &res, board)
	isAlive(i+1, j+1, &res, board)
	isAlive(i, j+1, &res, board)
	isAlive(i+1, j, &res, board)
	isAlive(i+1, j-1, &res, board)
	isAlive(i-1, j+1, &res, board)
	return res
}

func isAlive(i, j int, res *int, board [][]int) {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
		if board[i][j] == 1 || board[i][j] == 2 {
			*res = *res + 1
		}
	}
}
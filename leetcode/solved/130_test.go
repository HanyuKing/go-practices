package solved

import (
	"testing"
)

func TestSlove(t *testing.T) {
	board := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}

	println(solve(board))
	//
	//row := len(board)
	//col := len(board[0])
	//
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		fmt.Printf("%c ", board[i][j])
	//	}
	//	fmt.Printf("\n")
	//}

}

func solve(board [][]byte) int {
	row := len(board)
	col := len(board[0])

	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			count += visit(i, j, board, visited)
		}
	}
	return count
}

func visit(i, j int, board [][]byte, visited [][]bool) int {
	row := len(board)
	col := len(board[0])
	count := 0
	if visited[i][j] {
		return count
	}

	if board[i][j] == '1' {
		visited[i][j] = true
		count = 1
		if j < col-1 && board[i][j+1] == '1' {
			visit(i, j+1, board, visited)
		}
		if i < row-1 && board[i+1][j] == '1' {
			visit(i+1, j, board, visited)
		}
		if i > 0 && board[i-1][j] == '1' {
			visit(i-1, j, board, visited)
		}
		if j > 0 && board[i][j-1] == '1' {
			visit(i, j-1, board, visited)
		}
	}

	return count
}

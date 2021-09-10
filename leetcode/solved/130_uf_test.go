package solved

import (
	"fmt"
	"testing"
)

func TestUFSlove(t *testing.T) {
	board := [][]byte{
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'},
	}

	ufSolve(board)

	row := len(board)
	col := len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Printf("\n")
	}

}

func ufSolve(board [][]byte) {
	row, col := len(board), len(board[0])
	ufInit(row*col + 1)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				if board[i][j] == 'O' {
					ufMerge(i*col+j, row*col)
				}
			} else if board[i][j] == 'O' {
				currLocation := i*col + j
				if board[i][j-1] == 'O' { // left
					ufMerge(currLocation, i*col+(j-1))
				}
				if board[i-1][j] == 'O' { // up
					ufMerge(currLocation, (i-1)*col+j)
				}
				if board[i][j+1] == 'O' { // right
					ufMerge(currLocation, i*col+(j+1))
				}
				if board[i+1][j] == 'O' { // down
					ufMerge(currLocation, (i+1)*col+j)
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if ufFind(i*col+j) != ufFind(row*col) {
				board[i][j] = 'X'
			}
		}
	}

	return
}

var father []int
var rank []int

func ufInit(n int) {
	father = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		father[i] = i
		rank[i] = 1
	}
}

func ufFind(x int) int {
	if x == father[x] {
		return x
	} else {
		father[x] = ufFind(father[x])
		return father[x]
	}
}

func ufMerge(i, j int) {
	iF := ufFind(i)
	jF := ufFind(j)
	if rank[iF] < rank[jF] {
		father[iF] = jF
	} else if rank[iF] > rank[jF] {
		father[jF] = iF
	} else if rank[iF] == rank[jF] {
		father[iF] = jF
		if iF != jF {
			rank[jF]++
		}
	}
}

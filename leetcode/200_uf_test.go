package leetcode

import (
	"fmt"
	"testing"
)

func TestUFSlove(t *testing.T) {
	board := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	println(ufSolve(board))

	row := len(board)
	col := len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Printf("\n")
	}

}

func ufSolve(board [][]byte) int {
	row, col := len(board), len(board[0])
	ufInit(row*col + 1)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == '1' {
				currLocation := i*col + j
				if (j-1) >= 0 && board[i][j-1] == '1' { // left
					ufMerge(currLocation, i*col+(j-1))
				}
				if (i-1) >= 0 && board[i-1][j] == '1' { // up
					ufMerge(currLocation, (i-1)*col+j)
				}
				if (j+1) < col && board[i][j+1] == '1' { // right
					ufMerge(currLocation, i*col+(j+1))
				}
				if (i+1) < row && board[i+1][j] == '1' { // down
					ufMerge(currLocation, (i+1)*col+j)
				}
			}
		}
	}

	pMap := make(map[int]struct{})
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			p := ufFind(i*col + j)
			if p == i*col+j {
				if board[i][j] == '1' {
					pMap[p] = struct{}{}
				}
			} else {
				pMap[p] = struct{}{}
			}
		}
	}

	return len(pMap)
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

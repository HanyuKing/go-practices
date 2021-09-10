package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// 16:58
func TestMerge(t *testing.T) {
	intervals := make([][]int, 2)
	intervals[0] = []int{1, 4}
	intervals[1] = []int{0, 0}
	//intervals[2] = []int{8,10}
	//intervals[3] = []int{15,18}

	print(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	mergedIndex := make(map[int]int)
	newLen := len(intervals)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			minVal := min(intervals[i-1][0], intervals[i][0])
			canMerge := false
			if minVal == intervals[i][0] { // current
				if intervals[i][1] >= intervals[i-1][0] {
					canMerge = true
				}
			} else {
				if intervals[i-1][1] >= intervals[i][0] {
					canMerge = true
				}
			}
			if canMerge {
				intervals[i][0] = minVal
				intervals[i][1] = max(intervals[i-1][1], intervals[i][1])
				mergedIndex[i-1] = 1
				newLen--
			}
		}
	}
	newIntervals := make([][]int, newLen)
	newIndex := 0
	for i := 0; i < len(intervals); i++ {
		if mergedIndex[i] == 0 {
			newIntervals[newIndex] = intervals[i]
			newIndex++
		}
	}
	return newIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func print(intervals [][]int) {
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals[i]); j++ {
			fmt.Print(fmt.Sprintf("%d ", intervals[i][j]))
		}
		fmt.Println()
	}
}

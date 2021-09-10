package leetcode

import (
	"fmt"
	"testing"
)

func TestQSort(t *testing.T) {
	a := []int{4, 4, 23, 1, 4, 5, 9, 99, 100}
	qsort(a, 0, len(a)-1)
	for i := 0; i < len(a); i++ {
		fmt.Print(fmt.Sprintf("%d ", a[i]))
	}
	fmt.Println()
}

func qsort(a []int, i, j int) {
	low, high := i, j
	if low >= high {
		return
	}

	t := a[i]

	for low != high {

		for low < high {
			if t <= a[high] {
				high--
			} else {
				break
			}
		}

		if low < high {
			a[i] = a[high]
		}

		for low < high {
			if t >= a[low] {
				low++
			} else {
				break
			}
		}

		if low < high {
			a[high] = a[low]
		}

	}

	a[low] = t
	qsort(a, i, low-1)
	qsort(a, low+1, j)
}

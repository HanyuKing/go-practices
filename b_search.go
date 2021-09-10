package main

import "fmt"

// 1 3 5 7 9

// 1 2 3 4 5 6 7

func main() {
	// array := []int{1, 2, 3, 4, 5, 6, 7}
	array := []int{1, 3, 5, 7, 9}
	fmt.Println(bSearch(array, 6))
}

func bSearch(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low - 1
}

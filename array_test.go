package main

import (
	"fmt"
	"testing"
)

func Test_ArrayAsAParams(t *testing.T) {
	x := [3]int{1, 2, 3}

	func(arr [3]int) { // 数组是值拷贝，slice是地址
		arr[0] = 7
		fmt.Println(arr) // [7 2 3]
	}(x)
	fmt.Println(x) // [1 2 3]    // 并不是你以为的 [7 2 3]
}

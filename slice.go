package main

import "fmt"

// 包括下边界，不包括上边界
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	fmt.Println("///////////////////////////////////////////////////////////////")

	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)

	fmt.Println("///////////////////////////////////////////////////////////////")
	nilSlice()

	// Go中对nil的Slice和空Slice的处理是一致的吗
	// http://wearygods.online/articles/2021/04/19/1618823886966.html#toc_h4_28
	// panic: runtime error: index out of range [0] with length 0
	//
	//var slice []int
	//slice[1] = 0

	//slice := make([]int, 0)
	//slice[1] =  0
	//

	s = append(s, 1)
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

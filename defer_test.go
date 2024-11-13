package main

import (
	"fmt"
	"testing"
)

// defer is a stack

func Test_defer(t *testing.T) {
	defer fmt.Println("world")
	defer fmt.Println("hanyu")

	fmt.Println("hello")

	deferTest()
}

func deferTest() {
	fmt.Print("start\n")

	defer fmt.Print("\n")
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}

	fmt.Print("end\n")
}

// 对 defer 延迟执行的函数，它的参数会在声明时候就会求出具体值，而不是在执行时才求值：
func Test_defer2(t *testing.T) {
	var i = 2
	defer fmt.Println("result: ", func() int { return i * 2 }()) // 4
	i++
}

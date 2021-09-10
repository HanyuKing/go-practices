package main

import (
	"fmt"
	"strconv"
)

func strAppend(x int, y int) string {
	return strconv.Itoa(x) + " " + strconv.Itoa(y)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

var c, python, java bool

// 初始化变量
var initParams, j int = 1, 2

func main() {
	fmt.Print("append:", strAppend(1, 2), "\n")
	x, y := swap(1, 2)
	fmt.Print("swap: ", x, y, "\n")

	var i int
	fmt.Print(i, c, python, java, "\n")

	// 初始化变量
	fmt.Print("init params: ", initParams, j, "\n")
}

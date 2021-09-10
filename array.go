package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/**
	切片并不存储任何数据，它只是描述了底层数组中的一段。
	更改切片的元素会修改其底层数组中对应的元素。
	与它共享底层数组的切片都会观测到这些修改。
	*/
	// slice
	fmt.Println(primes[0:1])
}

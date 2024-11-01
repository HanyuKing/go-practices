package main

import (
	"fmt"
	"testing"
)

func Test_ArrayNil(t *testing.T) {
	var tmp []byte
	tmp = nil
	fmt.Println(len(tmp))

	a := []int{1, 2, 3}
	aa := a
	aa[0] = 11
	fmt.Printf("%p %d\n", &a, &a[0])
	fmt.Printf("%p %d\n", &aa, &aa[0])
	fmt.Println(a)

	var b = 1
	bb := b
	fmt.Println(&b)
	fmt.Println(&bb)
}

func Test_Append(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(len(x), cap(x))

	x = append(x, 6)
	fmt.Println(len(x), cap(x))

	x = append(x, 7, 8, 9, 10)
	y := x
	y[0] = 111
	fmt.Println(x, y, len(x), cap(x), len(y), cap(y))

	x = append(x, 11)
	y[0] = 11
	fmt.Println(x, y, len(x), cap(x), len(y), cap(y))
}

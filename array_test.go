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

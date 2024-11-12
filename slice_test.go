package main

import (
	"fmt"
	"testing"
)

func Test_SliceNil(t *testing.T) {
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

func Test_SliceCopy(t *testing.T) {
	f := func(Slice []int) {
		fmt.Println(Slice[0])
	}
	a := []int{1, 2, 3}

	f(a)

	a[0] = 111
	f(a)

	modifyF := func(Slice []int) {
		Slice[0] = 123
	}
	modifyF(a)
	fmt.Println(a[0])
}

func Test_SliceCopy2(t *testing.T) {
	s1 := []int{10, 20, 30}
	s2 := []int{40, 50, 60}
	s3 := s1
	copy(s1, s2)
	fmt.Println(s1, s2, s3)

	s1[0] = 1
	fmt.Println(s1, s2, s3)

	s2[0] = 11
	fmt.Println(s1, s2, s3)

	s3[0] = 111
	fmt.Println(s1, s2, s3)
}

func Test_SliceCompare(t *testing.T) {

}

// todo: https://golang.design/go-questions/slice/vs-array/

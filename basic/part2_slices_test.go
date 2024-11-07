package main

import (
	"fmt"
	"testing"
)

/**
https://medium.com/@ninucium/go-interview-questions-part-2-slices-87f5289fb7eb
*/

func Test4(t *testing.T) {
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	y := x
	x = append(x, 4)

	/*
		x:  1 2 3 4
		y:  1 2 3
	*/

	y = append(y, 5)
	x[0] = 0

	fmt.Println(x) // 0, 2, 3, 5
	fmt.Println(y) // 0, 2, 3, 5

	// What will the lines fmt.Println(x) and fmt.Println(y) output?
}

func Test5(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := x

	fmt.Printf("%d \n", &x[0])
	fmt.Printf("%d \n", &y[0])

	x = append(x, 5)
	y = append(y, 6)
	x[0] = 0

	fmt.Printf("%d %d \n", &x[0], &x[1])
	fmt.Printf("%d %d \n", &y[0], &y[1])

	fmt.Println(x) // 0,2,3,4,5
	fmt.Println(y) // 1,2,3,4,6
}

func Test6(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7) // 1,2,3,4,5,6,7
	a := x[4:]       // 5,6,7

	fmt.Printf("%d \n", &x[4])
	fmt.Printf("%d \n", &a[0])

	y := alterSlice(a)
	fmt.Println(x) // 1,2,3,4,10,6,7
	fmt.Println(y) // 10,6,7,11

	fmt.Println(x[0:8]) // 1,2,3,4,10,6,7,11
}

func alterSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

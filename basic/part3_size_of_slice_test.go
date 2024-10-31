package main

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
https://medium.com/@ninucium/go-interview-questions-part-3-size-of-slices-and-int-25080e51cf72
*/
func Test7(t *testing.T) {
	var emptySlice []byte
	sliceOfInts := []int{0, 1, 2, 3}
	fmt.Println(unsafe.Sizeof(emptySlice))  // 24
	fmt.Println(unsafe.Sizeof(sliceOfInts)) // 24

	sizeOfArray := unsafe.Sizeof(sliceOfInts[0]) * uintptr(len(sliceOfInts)) // sizeOfArray := uintptr(len(sliceOfInt)) * reflect.TypeOf(sliceOfInt).Elem().Size()
	fmt.Println(sizeOfArray)                                                 // 32

	var a = int(1)
	fmt.Printf("%T \n", a)
	fmt.Println(unsafe.Sizeof(a))

}

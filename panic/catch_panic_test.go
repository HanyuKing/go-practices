package main

import (
	"fmt"
	"testing"
)

func Test_CatchPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	fmt.Println("start...")

	count := []int{}
	fmt.Println(count[1])

	fmt.Println("1...")
}

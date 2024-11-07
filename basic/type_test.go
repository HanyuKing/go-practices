package main

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	var x, y, z = 9.6, 10, "codingninjas"
	fmt.Printf("%T %T %T \n", x, y, z) // float64 int string
}

package main

import (
	"fmt"
	"testing"
)

func Test_ArrayNil(t *testing.T) {
	var tmp []byte
	tmp = nil
	fmt.Println(len(tmp))
}

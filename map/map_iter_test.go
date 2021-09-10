package _map

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap_iter(t *testing.T) {
	mapData := make(map[string]interface{})
	mapData["a"] = 1
	mapData["b"] = 2
	for key := range mapData {
		fmt.Println(key)
	}

	arr := make([]uint64, 0)
	arr = append(arr, 1)
	fmt.Println(strings.Contains("asdasdas", ""))
}

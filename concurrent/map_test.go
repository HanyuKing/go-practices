package concurrent

import (
	"fmt"
	"testing"
)

func TestReadWrite(t *testing.T) {
	mapData := make(map[int]int)

	for i := 0; i < 1000; i++ {
		go func() {
			mapData[i] = i
		}()
	}

	for i := 0; i < 1000; i++ {
		fmt.Println(mapData[i])
	}
}

package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

// thread-unsafe
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

// thead-safe
func TestReadWriteSafe(t *testing.T) {
	mapData := sync.Map{}

	for i := 0; i < 1000; i++ {
		go func() {
			mapData.Store(i, i)
		}()
	}

	for i := 0; i < 1000; i++ {
		fmt.Println(mapData.Load(i))
	}
}

func TestMapCopy(t *testing.T) {
	map1 := map[int]int{
		1: 11,
		2: 22,
		3: 33,
	}
	map2 := map[int]int{}

	fmt.Println(map1)
	fmt.Println(map2)

	for k, v := range map1 {
		map2[k] = v
	}

	fmt.Println(map1)
	fmt.Println(map2)
}

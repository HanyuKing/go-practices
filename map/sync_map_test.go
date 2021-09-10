package _map

import (
	"fmt"
	"testing"
)

//func Test_SyncMapInit(t *testing.T) {
//	var mapA sync.Map
//	var mapB sync.Map
//	mapA.Store(1, 2)
//	mapB.Store(1, 3)
//
//	go func() {
//		for i := 0; i < 100000000 ; i++ {
//			var mapT sync.Map
//			mapT.Store(1, i)
//			mapA = mapT
//		}
//	}()
//
//	for i := 0; i < 1000000000; i++ {
//		val, ok := mapA.Load(1)
//		if !ok {
//			fmt.Println(val)
//		}
//	}
//}

func Test_MapInit(t *testing.T) {
	mapA := make(map[int]int)
	mapB := make(map[int]int)
	mapA[1] = 2
	mapB[1] = 3

	go func() {
		for i := 0; i < 100000000; i++ {
			mapT := make(map[int]int)
			mapT[1] = i
			mapA = mapT
		}
	}()

	for i := 0; i < 1000000000; i++ {
		val, ok := mapA[1]
		if !ok {
			fmt.Println(val)
		}
	}
}

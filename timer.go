package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		// do something

		now := time.Now()
		// 计算下一个触发点
		next := now.Add(24 * time.Hour)

		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		fmt.Println(123)
	}
}

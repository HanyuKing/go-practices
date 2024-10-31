package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		start := 1
		end := 10
		for start <= end {
			<-ch1
			fmt.Println(fmt.Sprintf("g1: %d", start))
			start++
			ch2 <- struct{}{}
		}
	}()

	go func() {
		start := 1
		end := 10
		for start <= end {
			<-ch2
			fmt.Println(fmt.Sprintf("g2: %d", start))
			start++
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	time.Sleep(time.Second)
}

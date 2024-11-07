package concurrent

import (
	"fmt"
	"testing"
)

func Test_Channel(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	count := 10
	go func() {
		currCount := 0
		for {
			select {
			case v := <-chan1:
				fmt.Printf("chan1: %d currCount: %d\n", v, currCount)
				chan2 <- v + 1
				currCount++
				if currCount == count {
					return
				}
			}
		}
	}()
	go func() {
		currCount := 0
		for {
			select {
			case v := <-chan2:
				fmt.Printf("chan2: %d currCount: %d\n", v, currCount)
				chan1 <- v + 1
				currCount++
				if currCount == count {
					return
				}
			}
		}
	}()
	chan1 <- 1
}

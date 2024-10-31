package algorithm

import (
	"fmt"
	"sync"
	"testing"
)

func getWorker(waitCh chan int, symbol int, wg *sync.WaitGroup) (next chan int) {
	notify := make(chan int)
	wg.Add(1)
	go func(waitCh chan int) {
		defer func() {
			wg.Done()
		}()
		for d := range waitCh {
			if d >= 30 {
				break
			}
			fmt.Println("goroutine:", symbol, "print", d+1)
			notify <- d + 1
		}
		close(notify)
		fmt.Println("goroutine: finish", symbol)
	}(waitCh)
	return notify
}

func TestPrint(t *testing.T) {
	wg := new(sync.WaitGroup)
	start := make(chan int)
	lastCh := start
	for i := 0; i < 3; i++ {
		lastCh = getWorker(lastCh, i+1, wg)
	}
	start <- 0
	for v := range lastCh {
		start <- v
	}
	close(start)
	wg.Wait()
}

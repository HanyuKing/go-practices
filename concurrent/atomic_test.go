package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var myCounter atomic.Int64
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myCounter.Add(1) // Each goroutine increments the myCounter by 1.
		}(i)
	}
	wg.Wait()
	fmt.Println(myCounter.Load()) // Display the final value of the counter.
}

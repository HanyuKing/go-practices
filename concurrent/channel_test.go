package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// Create a channel for synchronization between goroutines.
	ch := make(chan int)
	// Create a WaitGroup to wait for the goroutines to finish.
	var wg sync.WaitGroup
	// Goroutine 1.
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Perform some work in the first goroutine.
		fmt.Println("Goroutine 1.")
		time.Sleep(time.Second)
		ch <- 42
	}()
	// Goroutine 2.
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Perform some work in the second goroutine.
		fmt.Println("Goroutine 2.")
		result := <-ch
		fmt.Println("Result from Goroutine 1:", result)
	}()
	// Wait for both goroutines to finish.
	wg.Wait()
	// Close the channel.
	close(ch)
}

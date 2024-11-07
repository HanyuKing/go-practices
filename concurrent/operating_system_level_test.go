package concurrent

import (
	"fmt"
	"runtime"

	"sync"
	"testing"
)

/*
https://medium.com/@ninucium/parallelism-and-concurrency-in-go-how-it-works-in-real-computing-systems-part-2-b423288d047d
*/

func worker(id int, wg *sync.WaitGroup) {
	// Lock the current goroutine to an OS thread.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d: Iteration %d\n", id, i)
	}
}

func Test(t *testing.T) {
	// Create a wait group to synchronize goroutines.
	var wg sync.WaitGroup

	// Create and start three worker goroutines.
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for the worker goroutines to finish.
	wg.Wait()
}

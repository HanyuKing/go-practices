package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	myMap := map[int]int{} // Create a map to store integer values.
	mu := sync.RWMutex{}   // Create a read-write mutex for synchronization.
	wg := sync.WaitGroup{} // Create a wait group to wait for all goroutines to finish.
	for i := 0; i < 5; i++ {
		wg.Add(1) // Increment the wait group counter for each goroutine.
		go func(i int) {
			defer wg.Done() // Notify the wait group when the goroutine finishes.
			//mu.Lock()   // Lock the mutex for exclusive access during read and write operations.
			myMap[i] = i // Perform a write operation on the map.
			mu.Unlock()  // Unlock the mutex after the write operation is complete.
		}(i)
	}
	wg.Wait()          // Wait for all goroutines to finish before proceeding.
	fmt.Println(myMap) // Print the contents of the map.
}

package concurrent

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

// deadlock test
// 这是因为 wg 给拷贝传递到了 goroutine 中，导致只有 Add 操作，其实 Done操作是在 wg 的副本执行的
func TestWaitGroupDeadLock(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			fmt.Println(i)
		}(&wg, i) // 传引用，否则 deadlock
	}

	wg.Wait()

	fmt.Println("done...")
}

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()

	fmt.Println("done...")
}

func TestChannel(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		fmt.Println("start working")
		time.Sleep(time.Second * 1)
		ch <- struct{}{}
	}()

	<-ch

	fmt.Println("finished")
}

/*
*
https://medium.com/@ninucium/golang-concurrency-patterns-for-select-done-errgroup-and-worker-pool-645bec0bd3c9
*/
func TestForSelectDone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go PeriodicTask(ctx)

	// Create a channel to receive signals from the operating system.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	// The code blocks until a signal is received (e.g. Ctrl+C).
	<-sigCh
}
func someTask() {
	fmt.Println(rand.Int() * rand.Int())
}

// PeriodicTask runs someTask every 1 second.
// If canceled goroutine should be stopped.
func PeriodicTask(ctx context.Context) {
	// Create a new ticker with a period of 1 second.
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			someTask()
		case <-ctx.Done():
			fmt.Println("stopping PeriodicTask")
			ticker.Stop()
			return
		}
	}
}

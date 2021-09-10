package concurrent

import (
	"fmt"
	"sync"
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

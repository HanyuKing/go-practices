package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

// 等待所有 goroutine 执行完毕
// 进入死锁
/**
原因是 goroutine 得到的 "WaitGroup" 变量是 var wg WaitGroup 的一份拷贝值，即 doIt() 传参只传值。
所以哪怕在每个 goroutine 中都调用了 wg.Done()， 主程序中的 wg 变量并不会受到影响。
*/
func Test_DeadLock(t *testing.T) {
	var wg sync.WaitGroup
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done, wg) // 改这一行
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doIt(workerID int, done <-chan struct{}, wg sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}

func Test_UnDeadLock(t *testing.T) {
	var wg sync.WaitGroup
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt2(i, done, &wg)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doIt2(workerID int, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}

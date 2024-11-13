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

/*
*
向已关闭的 channel 发送数据会造成 panic

从已关闭的 channel 接收数据是安全的：接收状态值 ok 是 false 时表明 channel 中已没有数据可以接收了。
类似的，从有缓冲的 channel 中接收数据，
缓存的数据获取完再没有数据可取时，状态值也是 false

向已关闭的 channel 中发送数据会造成 panic：
*/
func Test_SendDataToClosedChanelPanic(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
			fmt.Println(fmt.Sprintf("send %d done", idx))
		}(i)
	}

	fmt.Println(<-ch)           // 输出第一个发送的值
	time.Sleep(2 * time.Second) // 模拟做其他的操作
	close(ch)                   // 不能关闭，还有其他的 sender

}

func Test_SendDataToClosedChanel(t *testing.T) {
	ch := make(chan int)
	done := make(chan interface{})
	for i := 0; i < 3; i++ {
		go func(idx int, done chan interface{}) {
			select {
			case ch <- idx:
				fmt.Println(fmt.Sprintf("send %d done", idx))
			case <-done:
				fmt.Println(idx, " channel closed!")
			}
		}(i, done)
	}

	fmt.Println(<-ch)           // 输出第一个发送的值
	time.Sleep(2 * time.Second) // 模拟做其他的操作
	close(done)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine的优雅退出方法有三种:

	// 1. 使用for-range退出
	//
	//in := make(chan int)
	//go func(in <-chan int) {
	//	// Using for-range to exit goroutine
	//	// range has the ability to detect the close/end of a channel
	//	for x := range in {
	//		fmt.Printf("Process %d\n", x)
	//	}
	//}(in)
	//
	//for i := 0; i < 10; i++ {
	//	in <- i
	//}

	// 2. 使用select case ,ok退出
	//in := make(chan int)
	//processedCnt := 0
	//go func() {
	//	// in for-select using ok to exit goroutine
	//	for {
	//		select {
	//		case x, ok := <-in:
	//			if !ok {
	//				fmt.Printf("No ok %d\n", x)
	//				return
	//			}
	//			fmt.Printf("Process %d\n", x)
	//			processedCnt++
	//		case <-time.After(1 * time.Second):
	//			fmt.Printf("Working, processedCnt = %d\n", processedCnt)
	//		}
	//	}
	//}()
	//in <- 1
	//in <- 2
	//time.Sleep(time.Second)
	//close(in)
	//time.Sleep(time.Second)

	// 多个通道关闭时
	//in1 := make(chan int)
	//in2 := make(chan int)
	//go func() {
	//	// in for-select using ok to exit goroutine
	//	for {
	//		select {
	//		case x, ok := <-in1:
	//			if !ok {
	//				in1 = nil
	//			}
	//			// Process
	//			fmt.Println(fmt.Sprintf("x: %d", x))
	//		case y, ok := <-in2:
	//			if !ok {
	//				in2 = nil
	//			}
	//			// Process
	//			fmt.Println(fmt.Sprintf("y: %d", y))
	//		case <-time.After(time.Second):
	//			fmt.Printf("Working. \n")
	//		}
	//
	//		// If both in channel are closed, goroutine exit
	//		if in1 == nil && in2 == nil {
	//			fmt.Printf("Exit. \n")
	//			return
	//		}
	//	}
	//}()
	//time.Sleep(5 * time.Second)
	//in1 <- 1
	//in2 <- 2
	//
	//close(in1)
	//time.Sleep(time.Second)
	//in2 <- 22
	//time.Sleep(time.Second)
	//close(in2)

	// 3. 使用退出通道
	//stopCh := make(chan struct{})
	//worker(stopCh)
	//time.Sleep(5 * time.Second)
	//stopCh <- struct{}{}
	//time.Sleep(1 * time.Second)
}

func worker(stopCh <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")
		// Using stop channel explicit exit
		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Working .")
			}
		}
	}()
	return
}

package main

import (
	"fmt"
	"sync"
	"testing"
)

/**
https://medium.com/@ninucium/go-interview-questions-part-1-pointers-channels-and-range-67c61345cf3c
*/
/*
func main() {
  ch := make(chan *int, 4)
  array := []int{1, 2, 3, 4}
  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    for _, value := range array {
      ch <- &value
    }
  }()
  go func() {
    for value := range ch {
      fmt.Println(*value) // what will be printed here?
    }
    wg.Done()
  }()

  wg.Wait()
}
*/
func Test1(t *testing.T) {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(len(array)) // 2
	go func() {
		for _, value := range array {
			t := value // 2
			ch <- &t
		}
		// close(ch) // 1
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value) // what will be printed here?
			wg.Done()
		}

	}()

	wg.Wait()
}

func Test2(t *testing.T) {
	array := []int{1, 2, 3, 4}
	newArray := make([]*int, 4)
	for i, value := range array {
		t := value // modify hear
		newArray[i] = &t
	}
	for _, value := range newArray {
		fmt.Println(*value)
	}
}

func Test3(t *testing.T) {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			ch <- &value
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value)
		}
		wg.Done()
	}()

	// New goroutine is run.
	go func() {
		var a int
		for {
			a++
		}
	}()

	wg.Wait()
}

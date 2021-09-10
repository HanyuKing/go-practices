package main

import (
	"fmt"
	"time"
)

func main() {
	ji := make(chan int)
	ou := make(chan int)
	i := 1

	go func() {
		for {
			num := <-ji

			fmt.Println(fmt.Sprintf("奇数: %d", num))

			time.Sleep(1 * time.Second)

			ou <- num + 1
		}
	}()

	go func() {
		for {
			num := <-ou

			fmt.Println(fmt.Sprintf("偶数: %d", num))

			time.Sleep(1 * time.Second)

			ji <- num + 1
		}
	}()

	ji <- i

	done := make(chan int)

	<-done
}

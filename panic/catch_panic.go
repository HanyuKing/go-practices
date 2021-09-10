package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	fmt.Println("start...")

	count := []int{}
	fmt.Println(count[1])

	fmt.Println("1...")
}

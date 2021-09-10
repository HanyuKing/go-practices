package main

import "fmt"

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

// while ?
func forLoop2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// dead loop
func forLoop3() {
	i := 0
	for {
		if i == 1 {
			break
		}
	}
}

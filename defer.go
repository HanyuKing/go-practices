package main

import "fmt"

// defer is a stack

func main() {
	defer fmt.Println("world\n")
	defer fmt.Print("hanyu\n")

	fmt.Println("hello")

	deferTest()
}

func deferTest() {
	fmt.Print("start\n")

	defer fmt.Print("\n")
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}

	fmt.Print("end\n")
}

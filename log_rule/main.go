package main

import "fmt"

func main() {
	ForRage()
	For()
}

func ForRage() {
	sum := 0
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}

func For() {
	sum := 0
	arr := []int{1, 2, 3, 4, 5}
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum)
}

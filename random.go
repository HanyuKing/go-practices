package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().Unix())
		fmt.Println(rand.Intn(10))
	}
}

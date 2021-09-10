package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	for cost := 10; cost <= 15; cost++ {
		startedAt := time.Now()
		bcrypt.GenerateFromPassword([]byte("password"), cost)
		duration := time.Since(startedAt)
		fmt.Printf("cost: %d, duration: %v\n", cost, duration)
	}
}

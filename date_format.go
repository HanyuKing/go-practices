package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf(t.Format("2006-01-02 15:04:05"))
}

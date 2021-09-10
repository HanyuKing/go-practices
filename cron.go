package main

import (
	"fmt"
	"gopkg.in/robfig/cron.v1"
)

func main() {
	c := cron.New()
	c.AddFunc("0 0/2 * * * *", func() {
		fmt.Println("every 3 seconds executing")
	})

	go c.Start()
	defer c.Stop()

	select {}
}

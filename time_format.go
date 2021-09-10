package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2020-04-20 15:52:19"
	loc, _ := time.LoadLocation("Local") // important
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err == nil {
		fmt.Println(err)
		fmt.Println(theTime)
		fmt.Println(theTime.Unix())
	}

}

package main

import (
	"fmt"
	"runtime"
	"time"
)

// Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// 没有条件的 switch 同 switch true 一样。
func noConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("Good morning")
	case t.Hour() < 17:
		fmt.Print("Good afternoon")
	}
}

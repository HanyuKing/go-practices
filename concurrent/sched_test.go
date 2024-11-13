package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func Test_1(t *testing.T) {
	go say("world")
	say("hello")
}

package thread

import (
	"fmt"
	"testing"
)

type FooBar struct {
	n  int
	c1 chan struct{}
	c2 chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:  n,
		c1: make(chan struct{}, 1),
		c2: make(chan struct{}),
	}
	fb.c1 <- struct{}{}
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.c1
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()

		fb.c2 <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.c2
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.c1 <- struct{}{}
	}
}

func Test(t *testing.T) {
	fooBar := NewFooBar(1)

	go fooBar.Bar(func() {
		fmt.Print("foo")
	})

	go fooBar.Foo(func() {
		fmt.Print("bar")
	})

}

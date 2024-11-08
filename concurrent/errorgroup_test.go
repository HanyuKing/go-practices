package concurrent

//
//import (
//	"errors"
//	"fmt"
//	"time"
//
//	"golang.org/x/sync/errgroup"
//)
//
//// errFailure some custom error.
//var errFailure = errors.New("some error")
//
//func main() {
//	// Create errgroup.
//	group := errgroup.Group{}
//
//	// Run first task.
//	group.Go(func() error {
//		time.Sleep(5 * time.Second)
//		fmt.Println("doing some work 1")
//		return nil
//	})
//
//	// Run second task.
//	group.Go(func() error {
//		fmt.Println("doing some work 2")
//		return nil
//	})
//
//	// Run third task.
//	group.Go(func() error {
//		fmt.Println("doing some work 3")
//		return errFailure
//	})
//
//	// Wait for all goroutines to complete.
//	if err := group.Wait(); err != nil {
//		fmt.Printf("errgroup tasks ended up with an error: %v\n", err)
//	} else {
//		fmt.Println("all works done successfully")
//	}
//}

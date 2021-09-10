package rand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_Seed(t *testing.T) {
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Float64())
	}
}

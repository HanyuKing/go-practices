package main

import (
	"fmt"
	"math"
)

func ifCondition(x int) string {
	if x < 10 {
		return "yes"
	} else {
		return "no"
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Print(ifCondition(10))
}

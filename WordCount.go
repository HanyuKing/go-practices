package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range strs {
		val, exists := m[v]
		if exists {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

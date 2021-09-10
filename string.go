package main

import "unicode/utf8"

func main() {
	str := "你好😁"
	count := utf8.RuneCountInString(str)
	println(count)
}

package main

import (
	"errors"
	"fmt"
	"testing"
)

func Test_SubstringTest(t *testing.T) {
	str := "你好啊啊啊啊"
	fmt.Println(SubString(str, 0, 1))
}

func Test_EncryptionPhone(t *testing.T) {
	str := "1839402340324"
	var newStr string
	for i, c := range str {
		if i >= 2 && i <= 5 {
			newStr += "*"
		} else {
			newStr += string(c)
		}
	}
	fmt.Println(newStr)
}

func SubString(str string, startIndex int, endIndex int) (string, error) {
	if startIndex < 0 || endIndex < 0 || startIndex > endIndex {
		return "", errors.New("invalid startIndex or endIndex")
	}
	newStr := ""
	index := 0
	for _, c := range str {
		if index >= startIndex && index <= endIndex {
			newStr += fmt.Sprintf("%c", c)
		}
		index++
	}
	return newStr, nil
}

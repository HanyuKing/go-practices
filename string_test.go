package main

import (
	"errors"
	"fmt"
	"testing"
)

/*
*
https://github.com/0voice/Introduction-to-Golang/blob/main/Golang%20%E6%96%B0%E6%89%8B%E5%8F%AF%E8%83%BD%E4%BC%9A%E8%B8%A9%E7%9A%84%2050%20%E4%B8%AA%E5%9D%91.md

range 得到的索引是字符值（Unicode point / rune）第一个字节的位置，与其他编程语言不同，这个索引并不直接是字符在字符串中的位置。

注意一个字符可能占多个 rune，比如法文单词 café 中的 é。操作特殊字符可使用norm 包。

for range 迭代会尝试将 string 翻译为 UTF8 文本，对任何无效的码点都直接使用 0XFFFD rune（�）UNicode 替代字符来表示。如果 string 中有任何非 UTF8 的数据，应将 string 保存为 byte slice 再进行操作。
*/
func Test_RangeString(t *testing.T) {
	data := "abc123"
	for i, v := range []byte(data) {
		fmt.Printf("%d, %c \n", i, v) // 正确
	}

	fmt.Println("------------------------")

	data = "A\xfe\x02\xff\x04"
	for i, v := range data {
		fmt.Printf("%d, %#x \n", i, v) // 0x41 0xfffd 0x2 0xfffd 0x4    // 错误
	}

	fmt.Println("------------------------")

	for i, v := range []byte(data) {
		fmt.Printf("%d, %#x \n ", i, v) // 0x41 0xfe 0x2 0xff 0x4    // 正确
	}
}

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

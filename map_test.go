package main

import (
	"fmt"
	"testing"
)

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func Test_Basic(t *testing.T) {
	m = make(map[string]Vertex2)
	m["key"] = Vertex2{
		40.68433, -74.39967,
	}
	obj := m["key"]
	obj.Lat = float64(1)
	fmt.Println(m["key"])
}

func Test_GetUnExistsKey(t *testing.T) {
	// 和其他编程语言类似，如果访问了 map 中不存在的 key 返回 null,
	// 但是 Go 则会返回元素对应数据类型的零值，比如 nil、'' 、false 和 0，取值操作总有值返回，
	// 故不能通过取出来的值来判断 key 是不是在 map 中。

	// 错误的 key 检测方式
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if v := x["twoo"]; v == "" {
		fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
	}

	// 正确示例
	x = map[string]string{"one": "2", "two": "", "three": "3"}
	if _, ok := x["two"]; !ok {
		fmt.Println("key two is no entry")
	}
}

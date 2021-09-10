package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v.X)
	//
	//p := &v
	//p.X = 5
	//fmt.Println(v.X)

	fmt.Println(v1, (*p).X, v2, v3)
}

package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_BasicRange(t *testing.T) {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}

// 在 range 迭代 slice、array、map 时通过更新引用来更新元素
func Test_UpdateObject(t *testing.T) {
	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10 // data 中原有元素是不会被修改的
	}
	fmt.Println("data: ", data) // data:  [1 2 3]

	for i, v := range data {
		data[i] = v * 10 // 如果要修改原有元素的值，应该使用索引直接访问
	}
	fmt.Println("data: ", data) // data:  [10 20 30]

	pointData := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range pointData {
		v.num *= 10 // 直接使用指针更新
	}
	fmt.Println(pointData[0], pointData[1], pointData[2]) // &{10} &{20} &{30}
}

/*
*
for 语句中的迭代变量在每次迭代中都会重用，即 for 中创建的闭包函数接收到的参数始终是同一个变量，在 goroutine 开始执行时都会得到同一个迭代值：
*/
func Test_ClosureIterator(t *testing.T) {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func() {
			fmt.Println(fmt.Sprintf("%p %s", &v, v))
		}()
	}
	time.Sleep(1 * time.Second)
	// 输出 three three three

	// 改进1
	fmt.Println("--------改进1--------")
	for _, v := range data {
		go func(vTemp string) {
			fmt.Println(fmt.Sprintf("%p %s", &vTemp, vTemp))
		}(v)
		//time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Second)

	// 改进2
	fmt.Println("--------改进2--------")
	for _, v := range data {
		vTemp := v
		go func() {
			fmt.Println(fmt.Sprintf("%p %s", &vTemp, vTemp))
		}()
		//time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Second)

	fmt.Println("--------迭代元素地址--------")
	data2 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data2 { // 此时迭代值 v 是三个元素值的地址，每次 v 指向的值不同
		fmt.Print(fmt.Sprintf("%p ", v))
		go v.print()
	}
	time.Sleep(1 * time.Second)
	// 输出 one two three
}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

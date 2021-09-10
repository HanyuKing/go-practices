package main

import (
	"fmt"
	"git.inke.cn/cxcenter/queen/server/library/utils/urldecode"
	"math"
	"reflect"
)

func main() {
	l := "MS4xMzQxMzYyMDU3NDQwNTNlLTMxNA=="
	f := urldecode.ParseFloat64(l)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	fmt.Println(math.MaxFloat64)

	fmt.Println(math.Pow(2, 127) * (math.Pow(2, 24) - 1) / math.Pow(2, 23))
}

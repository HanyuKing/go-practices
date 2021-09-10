package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type PrivateStruct struct {
	age int `json:"age"`
}

func Test_Main(t *testing.T) {
	ps := PrivateStruct{age: 11}
	o := reflect.ValueOf(ps)
	fmt.Println(o.FieldByName("age").Int())
}

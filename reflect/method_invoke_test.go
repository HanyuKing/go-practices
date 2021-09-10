package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_MethodInvoke(t *testing.T) {
	values := callReflect(&AAAHandler{}, "Serve", 1, 2)
	for _, val := range values {
		fmt.Println(val)
	}
}

func callReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		return nil
	} else {
		return v.Call(inputs)
	}

}

type AAAHandler struct{}

func (s *AAAHandler) Serve(a, b int) (interface{}, int) {
	return a + b, 0
}

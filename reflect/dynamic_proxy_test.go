package reflect

import (
	"github.com/traefik/yaegi/interp"
	"testing"
)

const src = `package foo
func Bar(s string) string { return s + "-Foo" }`

func main() {
	i := interp.New(interp.Options{})

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}

	bar := v.Interface().(func(string) string)

	r := bar("Kung")
	println(r)
}

func Test_AAA(t *testing.T) {
	main()
}

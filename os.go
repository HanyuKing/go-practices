package main

import "fmt"
import "os"

func main() {
	var JAVAHOME string
	JAVAHOME = os.Getenv("CONSUL_ADDR")
	fmt.Println(JAVAHOME)
}

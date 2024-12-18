package server

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Test_Server(t *testing.T) {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":8888", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

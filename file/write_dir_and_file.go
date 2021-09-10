package main

import (
	"os"
	"path/filepath"
)

func main() {
	dir, filename := filepath.Split("logs/test")
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0777); err != nil {
			println(err)
		}
	}
	var filePath = filepath.Join(dir, filename+".log")
	f, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	f.Write([]byte("a"))

}

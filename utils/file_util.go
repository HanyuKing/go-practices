package utils

import (
	"os"
)

func AppendFile(fileName string, str string) error {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	buf := []byte(str)
	_, err = fd.Write(buf)
	if err != nil {
		return err
	}
	return fd.Close()
}

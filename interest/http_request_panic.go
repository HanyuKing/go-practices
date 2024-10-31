package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	send("https://www.baidu.com", nil)
}

func send(uri string, param map[string]string) (content []byte, err error) {
	// Declare http client
	client := &http.Client{}
	client.Timeout = 100 * time.Millisecond

	body := url.Values{}
	body.Set("client_id", "180100031051")
	for k, v := range param {
		body.Set(k, v)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Referer", "http://m.mi.com/")

	cookie := &http.Cookie{
		Name:  "serviceToken",
		Value: "serviceToken",
	}

	ctx := context.WithValue(nil, "aaa", "val")
	ctx.Value("aaa")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(fmt.Sprintf("resp: %+v", *resp))

	req.AddCookie(cookie)

	return
}

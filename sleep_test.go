package main

import (
	"git.inke.cn/BackendPlatform/golang/logging"
	"testing"
	"time"
)

func Test_Sleep(t *testing.T) {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	time.Sleep(time.Duration(1) * time.Second)
	cost := time.Now().UnixNano()/int64(time.Millisecond) - start
	logging.Errorf("MoodTopicExposureFilter scard cost time:%d", cost)
}

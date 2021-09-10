package main

import (
	"fmt"
	"net/http"

	consulApi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

// 使用consul源码中的watch包监听服务变化
func main() {
	var (
		err    error
		params map[string]interface{}
		plan   *watch.Plan
		ch     chan int
	)
	ch = make(chan int, 1)

	params = make(map[string]interface{})
	params["type"] = "keyprefix"
	params["prefix"] = "queen/topicfeed/fe/"
	plan, err = watch.Parse(params)
	if err != nil {
		panic(err)
	}
	plan.Handler = func(index uint64, result interface{}) {
		if kvs, ok := result.(consulApi.KVPairs); ok {
			for _, kv := range kvs {
				if true {
					fmt.Println("index: " + fmt.Sprintf("%d", index) + ", modify index: " +
						fmt.Sprintf("%d", kv.ModifyIndex) + ", " + kv.Key + " : " + string(kv.Value) + "\n")
				}

			}

			ch <- 1
		}
	}
	go func() {
		// your consul agent addr
		if err = plan.Run("127.0.0.1:8500"); err != nil {
			panic(err)
		}
	}()
	go http.ListenAndServe(":8080", nil)
	go register()
	for {
		<-ch
		fmt.Printf("get change\n")
	}
}

func register() {
	var (
		err    error
		client *consulApi.Client
	)
	client, err = consulApi.NewClient(&consulApi.Config{Address: "127.0.0.1:8500"})
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceRegister(&consulApi.AgentServiceRegistration{
		ID:   "",
		Name: "test",
		Tags: []string{"SERVER"},
		Port: 8080,
		Check: &consulApi.AgentServiceCheck{
			HTTP: "",
		},
	})
	if err != nil {
		panic(err)
	}
}

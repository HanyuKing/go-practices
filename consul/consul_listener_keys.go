package main

import (
	"fmt"
	"net/http"

	consulApi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

// 使用consul源码中的watch包监听服务变化
func main() {

	ch := make(chan int, 1)

	keys := []string{"queen/topicfeed/fe/mood_max_len", "queen/topicfeed/fe/mood_voice_max_len"}
	var plans []watch.Plan

	for _, key := range keys {
		params := make(map[string]interface{})
		params["type"] = "key"
		params["key"] = key
		plan, err := watch.Parse(params)
		if err != nil {
			panic(err)
		}

		plans = append(plans, *plan)
	}

	for _, plan := range plans {
		fmt.Println(plan.Watcher)
		plan.Handler = func(index uint64, result interface{}) {
			kv, ok := result.(*consulApi.KVPair)
			if ok {

				fmt.Println(kv.Key + " : " + string(kv.Value))

				ch <- 1
			}
		}

		plan.Run("127.0.0.1:8500")
	}

	go http.ListenAndServe(":8080", nil)
	go register2()
	for {
		<-ch
		fmt.Printf("get change\n")
	}
}

func register2() {
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
